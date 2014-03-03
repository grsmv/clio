package cli

import (
	"bytes"
	"fmt"
	"github.com/grsmv/clio/helpers"
	"github.com/grsmv/inflect"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Resource struct {
	PluralTitle, SingularTitle,
	PluralPath, SingularPath,
	AppName string
}

var templatesPaths = map[string]string{
	"controller":      "app/controllers/resources.go.tmpl",
	"router":          "app/routes/resources.go.tmpl",
	"view-index":      "app/views/resources/index.template.tmpl",
	"view-resource":   "app/views/resources/resource.template.tmpl",
	"controller-spec": "spec/controllers/resource.go.tmpl",
}

// todo: check if this operation executes in app's root
func NewResource(name string) Resource {
	wd, _ := os.Getwd()
	splittedPath := strings.Split(wd, slash)

	return Resource{
		PluralTitle: inflect.Camelize(
			inflect.Pluralize(name),
		),
		SingularTitle: inflect.Camelize(
			inflect.Singularize(name),
		),
		PluralPath: inflect.Underscore(
			strings.ToLower(
				inflect.Pluralize(name),
			),
		),
		SingularPath: inflect.Underscore(
			strings.ToLower(
				inflect.Singularize(name),
			),
		),
		AppName: splittedPath[len(splittedPath)-1],
	}
}

func (resource *Resource) Scaffold() error {
	resource.Router()
	resource.Controller()
	resource.View()
	resource.ControllerSpec()
	return nil
}

func (resource *Resource) Router() {
	resource.templatize(
		map[string]string{
			"router": "app/routes/" + resource.PluralPath + ".go",
		},
	)
}

func (resource *Resource) Controller() {
	resource.templatize(
		map[string]string{
			"controller": "app/controllers/" + resource.PluralPath + ".go",
		},
	)
}

func (resource *Resource) View() {

	parentFolder := "app/views/" + resource.PluralPath

	if _, err := os.Stat(parentFolder); os.IsNotExist(err) {
		os.Mkdir(parentFolder, 0755)
	}

	resource.templatize(
		map[string]string{
			"view-index":    parentFolder + "/index.template",
			"view-resource": parentFolder + "/" + resource.SingularPath + ".template",
		},
	)
}

func (resource *Resource) ControllerSpec() {
	resource.templatize(
		map[string]string{
			"controller-spec": "spec/controllers/" + resource.PluralPath + "_test.go",
		},
	)
}

func (resource *Resource) templatize(files map[string]string) error {

	// nb: generatorsTemplatesPath
	for templateType, file := range files {
		var buffer bytes.Buffer

		// preventing overwriting
		if _, err := os.Stat(file); err == nil {
			fmt.Println(yellow, "  exists:", reset, file)
			continue
		}

		// reading and processing template
		templateContents, err := ioutil.ReadFile(
			helpers.FixPath(
				generatorsTemplatesPath +
					string(os.PathSeparator) +
					templatesPaths[templateType],
			),
		)

		tmpl, err := template.New("generator").Parse(string(templateContents))
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buffer, resource)

		// writing template in appropriate folder
		err = ioutil.WriteFile(helpers.FixPath(file), buffer.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(green, "  create:", reset, file)
	}
	return nil
}

// vim: noai:ts=4:sw=4
