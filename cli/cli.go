package cli

import (
    "fmt"
    "os"
    "strings"
    "github.com/cliohq/clio/helpers"
    "github.com/cliohq/clio/core"
)

const (
    tcpIPCPort = 31000

    red    = "\x1b[0;31m"
    green  = "\x1b[0;32m"
    yellow = "\x1b[0;33m"
    reset  = "\x1b[0m"
)

var (
    slash = string (os.PathSeparator)

    GOPATH = os.Getenv ("GOPATH")
    GOPATH_SRC = GOPATH + slash + "src" + slash

    templatesRoot = "src/github.com/cliohq/clio/templates"
    applicationTemplatesPath = helpers.FixPath (templatesRoot + "/application")
    generatorsTemplatesPath  = helpers.FixPath (helpers.FixPath(GOPATH + "/" + templatesRoot + "/generators"))
)


/**
 *  Dispatching command-line arguments to a separate
 *  command execution
 */
func Route () {
    if len (os.Args) > 1 {

        switch os.Args[1] {

        case "help":
            Help()

        case "version":
            fmt.Println (core.Version ())

        case "run":
            Run()

        case "watch":
            Watch ()

        case "create":
            {
                if len(os.Args) == 3 {
                    Create(os.Args[2])
                } else {
                    fmt.Println(red, "Please specify application name", reset)
                }
            }

        case "generate":
        case "g":
            {
                if len (os.Args) > 3 {
                    resource := NewResource (strings.Join(os.Args[3:], " "))
                    switch os.Args[2] {
                        case "scaffold":
                        case "s":
                            resource.Scaffold ()

                        case "router":
                        case "r":
                            resource.Router ()

                        case "controller":
                        case "c":
                            resource.Controller ()

                        case "view":
                        case "v":
                            resource.View ()
                    }
                } else {
                    Help()
                }
            }

        default:
            {
                Help ()
            }
        }
    } else {
        Help()
    }
}

// vim: noai:ts=4:sw=4
