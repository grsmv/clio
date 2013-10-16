package cli

import (
    "fmt"
    "os"
)

const (
    red    = "\x1b[0;31m"
    green  = "\x1b[0;32m"
    yellow = "\x1b[0;33m"
    reset  = "\x1b[0m"
)

const tcpIPCPort = 31000

type Resource struct {
    name string
    fields []string
}

/**
 *  Dispatching command-line arguments to a separate
 *  command execution
 */
func Route () {
    if len (os.Args) > 1 {

        switch os.Args[1] {

        case "help":
            Help()
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

        case "g":
            {
                if len (os.Args) > 3 {
                    resource := Resource { name: os.Args[3], fields: os.Args[4:] }
                    switch os.Args[2] {
                    case "scaffold":
                        resource.GenerateScaffold ()
                    case "router":
                        resource.GenerateRouter ()
                    case "controller":
                        resource.GenerateController ()
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
