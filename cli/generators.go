package cli

import (
    "github.com/grsmv/inflect"
    "strings"
)


type Resource struct {
    pluralTitle, singularTitle,
    pluralPath, singularPath string
}


func NewResource (name string) Resource {
    return Resource {
        pluralTitle:   inflect.Camelize (
                          inflect.Pluralize (name),
                       ),
        singularTitle: inflect.Camelize (
                           inflect.Singularize (name),
                       ),
        pluralPath:    inflect.Underscore (
                           strings.ToLower (
                               inflect.Pluralize (name),
                           ),
                       ),
        singularPath:  inflect.Underscore (
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


func (resource *Resource) Router () error {
    // todo: a. Create contents of file from template
    //       b. Create and write file in app/routers/NAME
    return nil
}


func (resource *Resource) Controller () error {
    // todo: a. Create contents of file from template
    //       b. Create and write file in app/controllers/NAME
    return nil
}


func (resource *Resource) View () error {
    // todo: a. Create contents of files from template
    //       b. Create and write file in app/views/name/{index,name}.template
    return nil
}

// vim: noai:ts=4:sw=4
