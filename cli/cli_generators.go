package cli

import (
    "fmt"
)

func (resource *Resource) GenerateScaffold () {
    fmt.Println("Generate Backbone & Carcass scaffold with name", resource.name)
}

func (resource *Resource) GenerateRouter () {
    fmt.Println("Generate Backbone & Carcass router with name", resource.name)
}

func (resource *Resource) GenerateModel () {
    fmt.Println("Generate Backbone & Carcass model with name", resource.name)
    fmt.Println("model fields:", resource.fields)
}

func (resource *Resource) GenerateController () {
    fmt.Println("Generate Carcass controller with name", resource.name)
}

// vim: noai:ts=4:sw=4
