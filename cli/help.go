package cli

import (
    "fmt"
)

func Help() {
    fmt.Println (
        "                                                                 \n" +
        "Usage:                                                           \n" +
        "                                                                 \n" +
        "   clio COMMAND [FLAGS]                                          \n" +
        "                                                                 \n" +
        "Commands:                                                        \n" +
        "                                                                 \n" +
        "   create      Creating application skeleton                     \n" +
        "   g           Generating scaffold, controller, router or view   \n" +
        "   run         Running application and all helping workers       \n" +
        "   help        Output this message again                         \n" +
        "   version     Displaying version information                    \n" +
        "                                                                 \n" +
        "Generators:                                                      \n" +
        "                                                                 \n" +
        "   clio g [controller | view | router | scaffold] NAME           \n" +
        "                                                                 \n" +
        "Example:                                                         \n" +
        "                                                                 \n" +
        "   clio create awesome-application                               \n" +
        "                                                                 \n" +
        "   This will generate a skeletal Clio application.               \n" +
        "   Please, see the README in the newly created application       \n" +
        "   get going. Also http://cliohq.github.io is a good place       \n" +
        "   for getting more information about Clio. Good luck!           \n" +
        "                                                                 \n")
}


// vim: noai:ts=4:sw=4
