package cli

import (
    "bytes"
    "fmt"
    "github.com/grsmv/inflect"
    "io/ioutil"
    "log"
    "os"
    "strings"
    "text/template"
)


type Resource struct {
    PluralTitle, SingularTitle,
    PluralPath,  SingularPath  string
}


func NewResource (name string) Resource {
    return Resource {
        PluralTitle:   inflect.Camelize (
                          inflect.Pluralize (name),
                       ),
        SingularTitle: inflect.Camelize (
                           inflect.Singularize (name),
                       ),
        PluralPath:    inflect.Underscore (
                           strings.ToLower (
                               inflect.Pluralize (name),
                           ),
                       ),
        SingularPath:  inflect.Underscore (
                           strings.ToLower (
                               inflect.Singularize (name),
                           ),
                       ),
    }
}


func (resource *Resource) Scaffold () error {
    resource.Router ()
    resource.Controller ()
    resource.View ()
    return nil
}


func (resource *Resource) Router () {
    resource.templatize (
        []string { "app/routes/" + resource.PluralPath + ".go" },
    )
}


func (resource *Resource) Controller () {
    resource.templatize (
        []string { "app/controllers/" + resource.PluralPath + ".go" },
    )
}


func (resource *Resource) View () {

    parentFolder := "app/views/" + resource.PluralPath

    if _, err := os.Stat (parentFolder); os.IsNotExist (err) {
        os.Mkdir (parentFolder, 0755)
    }

    resource.templatize (
        []string {
            parentFolder + "/index.template",
            parentFolder + "/" + resource.SingularPath + ".template",
        },
    )
}


func (resource *Resource) templatize (files []string) error {

    // nb: generatorsTemplatesPath
    for _, file := range files {
        var buffer bytes.Buffer

        // preventing overwriting
        if _, err := os.Stat (file); err == nil {
            fmt.Println(yellow, "  exists:", reset, file)
            continue
        }

        // reading and processing template
        ps := string (os.PathSeparator)
        templateSource := strings.Replace (file, ps + resource.PluralPath + ps, ps + "resources" + ps, 1)

        templateContents, _ := ioutil.ReadFile (generatorsTemplatesPath + string (os.PathSeparator) + templateSource + ".tmpl")
        tmpl, err := template.New ("generator").Parse (string(templateContents)); if err != nil {
            log.Fatal (err)
        }
        tmpl.Execute (&buffer, resource)

        // writing template in appropriate folder
        err = ioutil.WriteFile (file, buffer.Bytes (), 0644); if err != nil {
            log.Fatal (err)
        }

        fmt.Println(green, "  create:", reset, file)
    }
    return nil
}


// vim: noai:ts=4:sw=4
