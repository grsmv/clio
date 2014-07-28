package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	red    = "\x1b[0;31m"
	green  = "\x1b[0;32m"
	yellow = "\x1b[0;33m"
	reset  = "\x1b[0m"
)

var (
	GOPATH                   = os.Getenv("GOPATH")
	GOPATH_SRC               = filepath.Join(GOPATH, "src")
	templatesRoot            = filepath.Join("src", "github.com", "grsmv", "clio", "templates")
	applicationTemplatesPath = filepath.Join(templatesRoot, "application")
	generatorsTemplatesPath  = filepath.Join(GOPATH, templatesRoot, "generators")
)

/**
 *  Dispatching command-line arguments to a separate
 *  command execution
 */
func Route() {
	if len(os.Args) > 1 {

		switch os.Args[1] {

		case "help":
			Help()

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
				if len(os.Args) > 3 {
					resource := NewResource(strings.Join(os.Args[3:], " "))
					switch os.Args[2] {
					case "scaffold":
					case "s":
						resource.Scaffold()

					case "router":
					case "r":
						resource.Router()

					case "controller":
					case "c":
						resource.Controller()

					case "view":
					case "v":
						resource.View()
					}
				} else {
					Help()
				}
			}

		default:
			{
				Help()
			}
		}
	} else {
		Help()
	}
}
