package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

type Settings struct {
	Layout string
}

var (
	templateExtension = ".html"
	templatesHolder   = filepath.Join("app", "views")
	defaultLayoutName = "application"
)

/**
 *  Generic function for rendering templates and partials
 *    @param <String> _type - template or partial
 *    @param <String> name - name (wothout extension) to lookup
 *    @param <Slice> obj - data structure to compile with template (can be empty)
 */
func templateGeneric(_type, name string, data []interface{}) (contents string) {
	var fileName = filepath.Join(templatesHolder, name+templateExtension)

	// checking if template file exists
	if _, err := os.Stat(fileName); err == nil {
		fileContents, _ := ioutil.ReadFile(fileName)
		contents = processTemplate(string(fileContents), data)
	} else {
		log.Fatal("ERROR: No such " + _type + ": " + fileName)
	}
	return
}

/**
 *  Executing template
 */
func processTemplate(unrocessedContents string, data []interface{}) string {
	var (
		// functions available inside templates
		customFunctions = template.FuncMap{
			"partial":     partial,
			"development": Development,
			"js":          IncludeJavascripts,
			"css":         IncludeStylesheets,
		}

		// defining source for template variables by default.
		// it's make possible to call render methods in templates
		// with and without data structure as additional argument
		dataStruct interface{}
		buffer     bytes.Buffer
	)

	// initializing new template
	rawTemplate := template.New("tmpl")

	// adding functions to be used inside templates
	rawTemplate = rawTemplate.Funcs(customFunctions)

	tmpl, err := rawTemplate.Parse(unrocessedContents)
	if err != nil {
		log.Fatal(err)
	}

	// redefining source for template variables if one given
	if len(data) > 0 {
		dataStruct = data[0]
	}

	// compiling and returning template's contents
	err = tmpl.Execute(&buffer, dataStruct)
	if err != nil {
		log.Fatal(err)
	}

	return buffer.String()
}

/**
 *  Returning wrapping layout for rendered template.
 *  Used in application files as:
 *    `Render("a/b", dataInterface, Settings { Layout: 'name' })`
 */
func layout(layoutName, renderedTemplate string) (output string) {
	if layoutName != "none" {
		layoutFilepath := filepath.Join(templatesHolder, "layouts", layoutName+templateExtension)
		layoutContents, err := ioutil.ReadFile(layoutFilepath)

		if err != nil {
			log.Fatal(err)
		} else {
			pattern := regexp.MustCompile("{{[\\s]*?yield[\\s]*?}}")
			output = pattern.ReplaceAllString(string(layoutContents), renderedTemplate)
		}
	} else {
		output = renderedTemplate
	}

	output = processTemplate(output, []interface{}{})
	return
}

/**
 *  Rendering basic template
 */
func Render(name string, obj ...interface{}) string {
	var layoutName = defaultLayoutName

	// wrapping rendered template by layout
	if len(obj) > 0 {
		for _, arg := range obj {
			settings := reflect.ValueOf(arg)
			if settings.Type().String() == "clio.Settings" {
				layoutName = settings.FieldByName("Layout").String()
				break
			}
		}
	}

	return layout(layoutName, templateGeneric("template", name, obj))
}

// template built-in helpers: ---------------------------------------

/**
 *  Rendering partial (used inside other templates as
 *    `{{ partial "partial_name"}}`)
 */
func partial(name string, obj ...interface{}) string {
	return templateGeneric("partial", name, obj)
}

func IncludeJavascripts(files ...string) string {
	return genericInclude("<script type=\"text/javascript\" src=\"/assets/{file}.js\"></script>", files)
}

func IncludeStylesheets(files ...string) string {
	return genericInclude("<link rel=\"stylesheet\" href=\"/assets/{file}.css\" type=\"text/css\" media=\"screen\" charset=\"utf-8\">", files)
}

func genericInclude(template string, files []string) (output string) {
	for _, file := range files {
		output = fmt.Sprintln(output, strings.Replace(template, "{file}", file, -1))
	}
	return
}
