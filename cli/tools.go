package cli

import (
    "fmt"
    "os"
)

func Help() {
    bin := os.Args[0]
    fmt.Println (
        "This program is a command line interface to the Clio framework infrastructure.         \n" +
        "See https://github.com/grsmv/clio/ for more information.                               \n" +
        "                                                                                       \n" +
        "Usage: ", bin, "COMMAND [FLAGS]                                                        \n" +
        "                                                                                       \n" +
        "Commads:                                                                               \n" +
        "   create      Creating application skeleton                                           \n" +
        "   build       Building application binary                                             \n" +
        "   g           Generating scaffold, model, router or view                              \n" +
        "   run         Running application and all helping workers                             \n" +
        "   stop        Killing all application's processes                                     \n" +
        "   test        Testing current application                                             \n" +
        "   clean       Removing temporary files and binary                                     \n" +
        "   deps        Installing application's dependencies, declared in config/dependencies  \n" +
        "   release     Deploy aplication, according to selected strategy                       \n" +
        "   help        Output this message again                                               \n" +
        "                                                                                       \n" +
        "Generators:                                                                            \n" +
        "   " + bin + " g [model | view | router | scaffold] NAME [ARGUMENTS]                   \n")
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
