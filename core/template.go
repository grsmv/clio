package core

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"text/template"
)

type Settings struct {
	Layout string
}

var (
	templateExtension = ".template"
	templatesHolder   = "app/views"
	defaultLayoutName = "application"
	sep               = string(os.PathSeparator)
)

/**
 *  Generic function for rendering templates and partials
 *    @param <String> _type - template or partial
 *    @param <String> name - name (wothout extension) to lookup
 *    @param <Slice> obj - data structure to compile with template (can be empty)
 */
func templateGeneric(_type, name string, obj []interface{}) (contents string) {
	var fileName = templatesHolder + sep + name + templateExtension

	// checking if template file exists
	if _, err := os.Stat(fileName); err == nil {

		fileContents, _ := ioutil.ReadFile(fileName)
		contents = processTemplate(string(fileContents), obj)

	} else {
		log.Fatal("ERROR: No such " + _type + ": " + fileName)
	}

	return
}

/**
 *  Executing template
 */
func processTemplate(unrocessedContents string, data ...interface{}) string {
	var (
		// functions available inside templates
		customFunctions = template.FuncMap{
			"partial":     partial,
			"development": Development,
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
		layoutFilepath := templatesHolder + sep + "layouts" + sep + layoutName + templateExtension
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

	output = processTemplate(output)
	return
}

/**
 *  Rendering partial (used inside other templates as
 *    `{{ partial "partial_name"}}`)
 */
func partial(name string, obj ...interface{}) string {
	return templateGeneric("partial", name, obj)
}

/**
 *  Displaying development state in templates with:
 *    `{{ development }}`
 */
func developmentState() bool {
	return development
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

// vim: noai:ts=4:sw=4
