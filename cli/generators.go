package cli

import (
    "fmt"
)

func (resource *Resource) GenerateScaffold () {
    fmt.Println ("Scaffold :", resource.name)
}

func (resource *Resource) GenerateRouter () {
    fmt.Println ("Router :", resource.name)
}

func (resource *Resource) GenerateController () {
    fmt.Println ("Controller :", resource.name)
}

// vim: noai:ts=4:sw=4
