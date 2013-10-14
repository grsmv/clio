package cli

import (
    "fmt"
)

func Help() {
    fmt.Println (
        "Usage:                                                        \n" +
        "   clio COMMAND [FLAGS]                                       \n" +
        "                                                              \n" +
        "Commads:                                                      \n" +
        "   create      Creating application skeleton                  \n" +
        "   g           Generating scaffold, model, router or view     \n" +
        "   run         Running application and all helping workers    \n" +
        "   help        Output this message again                      \n" +
        "                                                              \n" +
        "Generators:                                                   \n" +
        "   clio g [controller | view | router | scaffold] NAME        \n")
}


// vim: noai:ts=4:sw=4
