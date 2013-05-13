package clio

import (
    "bytes"
    "log"
    "os"
    "text/template"
)

var templateExtension = ".template"

/**
 *  Generic function for rendering templates and partials
 *    @param <String> _type - template or partial
 *    @param <String> name - name (wothout extension) to lookup
 *    @param <Slice> obj - data structure to compile with template (can be empty)
 */
func Template (_type, name string, obj []interface{}) (contents string) {
    var (
        // functions available inside templates
        customFunctions = template.FuncMap{ "partial": Partial }
        fileName = name + templateExtension

        // defining source for template variables by default.
        // it's make possible to call render methods in templates
        // with and without data structure as additional argument
        dataStruct interface{}
        buffer bytes.Buffer
    )

    // checking if template file exists
    if _, err := os.Stat(fileName); err == nil {

        // initializing new template
        rawTemplate := template.New(fileName)

        // adding functions to be used inside templates
        rawTemplate = rawTemplate.Funcs(customFunctions)
        tmpl, _ := rawTemplate.ParseFiles(fileName)

        // redefining source for template variables if one given
        if (len(obj) > 0) { dataStruct = obj[0] }

        // compiling and returning template's contents
        tmpl.Execute(&buffer, dataStruct)
        contents = buffer.String()

    } else {
        log.Fatal("ERROR: No such " + _type + ": " + fileName)
    }

    return
}


/**
 *  Rendering basic template
 */
func Render (name string, obj ...interface{}) string {
    return Template("template", name, obj)
}


/**
 *  Rendering partial (used inside other templates as
 *    `{{ partial "partial_name"}}`)
 */
func Partial (name string, obj ...interface{}) string {
    return Template("partial", name, obj)
}


// vim: noai:ts=4:sw=4
