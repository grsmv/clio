package cli

import (
    "github.com/grsmv/inflect"
    "strings"
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
    // todo: create folder `app/views/resources` if not exists
    resource.templatize (
        []string {
            "app/views/" + resource.PluralPath + "/index.template",
            "app/views/" + resource.PluralPath + "/" + resource.SingularPath + ".template",
        },
    )
}


func (resource *Resource) templatize (files []string) error {
    // todo: a. Read and process template
    //       b. Write template in appropriate folder
    for _, file := range files {
        println (file)
    }
    return nil
}

// vim: noai:ts=4:sw=4
