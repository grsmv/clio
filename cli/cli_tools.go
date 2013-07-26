package cli

import (
    "fmt"
    "os"
)

func Help() {
    fmt.Println("Usage:", os.Args[0], "<action>")
}

func Build() {
    fmt.Println(green, "app:", reset, "build")
}

func Clean() {
    fmt.Println(green, "app:", reset, "clean")
}

func Dependencies() {
    fmt.Println(green, "app:", reset, "install dependencies")
}

func Release() {
    fmt.Println(green, "app:", reset, "release")
}

func Stop() {
    fmt.Println(green, "app:", reset, "stop")
}

func Test() {
    fmt.Println(yellow, "test", reset)
}

// vim: noai:ts=4:sw=4
