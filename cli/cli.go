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

type Resource struct {
    name string
    fields []string
}

/**
 *  Dispatching command-line arguments to a separate
 *  command execution
 */
func Route() {
    if len(os.Args) > 1 {

        switch os.Args[1] {

        case "help":
            Help()
        case "build":
            Build()
        case "clean":
            Clean()
        case "deps":
            Dependencies()
        case "release":
            Release()
        case "run":
            Run()
        case "stop":
            Stop()
        case "test":
            Test()

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
                if len(os.Args) > 3 {
                    resource := Resource { name: os.Args[3], fields: os.Args[4:] }
                    switch os.Args[2] {
                    case "scaffold":
                        resource.GenerateScaffold ()
                    case "router":
                        resource.GenerateRouter ()
                    case "model":
                        resource.GenerateModel ()
                    case "controller":
                        resource.GenerateController ()
                    }
                } else {
                    Help()
                }
            }

        default:
            {
                fmt.Println("action:", os.Args[1])

                // current working directory
                wd, _ := os.Getwd()
                fmt.Println("working directory:", wd)
            }
        }
    } else {
        Help()
    }
}

// vim: noai:ts=4:sw=4
