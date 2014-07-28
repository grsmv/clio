package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/grsmv/inflect"
)

type Resource struct {
	PluralTitle, SingularTitle,
	PluralPath, SingularPath,
	AppName string
}

var templatesPaths = map[string]string{
	"controller":      filepath.Join("app", "controllers", "resources.go.tmpl"),
	"router":          filepath.Join("app", "routes", "resources.go.tmpl"),
	"view-index":      filepath.Join("app", "views", "resources", "index.html.tmpl"),
	"view-resource":   filepath.Join("app", "views", "resources", "resource.html.tmpl"),
	"controller-spec": filepath.Join("spec", "controllers", "resource.go.tmpl"),
}

// todo: check if this operation executes in app's root
func NewResource(name string) Resource {
	wd, _ := os.Getwd()
	splittedPath := strings.Split(wd, string(os.PathSeparator))

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
			"router": filepath.Join("app", "routes", resource.PluralPath+".go"),
		},
	)
}

func (resource *Resource) Controller() {
	resource.templatize(
		map[string]string{
			"controller": filepath.Join("app", "controllers", resource.PluralPath+".go"),
		},
	)
}

func (resource *Resource) View() {

	parentFolder := filepath.Join("app", "views", resource.PluralPath)

	if _, err := os.Stat(parentFolder); os.IsNotExist(err) {
		os.Mkdir(parentFolder, 0755)
	}

	resource.templatize(
		map[string]string{
			"view-index":    filepath.Join(parentFolder, "index.html"),
			"view-resource": filepath.Join(parentFolder, resource.SingularPath+".html"),
		},
	)
}

func (resource *Resource) ControllerSpec() {
	resource.templatize(
		map[string]string{
			"controller-spec": filepath.Join("spec", "controllers", resource.PluralPath+"_test.go"),
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
			filepath.Join(generatorsTemplatesPath, templatesPaths[templateType]),
		)

		tmpl, err := template.New("generator").Parse(string(templateContents))
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buffer, resource)

		// writing template in appropriate folder
		err = ioutil.WriteFile(file, buffer.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(green, "  create:", reset, file)
	}
	return nil
}
