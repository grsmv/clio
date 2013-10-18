package cli

import (
    "fmt"
    "os"
    "strings"
)

const (
    red    = "\x1b[0;31m"
    green  = "\x1b[0;32m"
    yellow = "\x1b[0;33m"
    reset  = "\x1b[0m"
)

const tcpIPCPort = 31000

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
