package cli

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

const templatesPath = "src/github.com/cliohq/clio/templates/application"

type Application struct {
    name string
}


/**
 *  High level abstraction for creating new app
 *  (called from cli.Route())
 */
func Create (appName string) {
    app := Application { name: appName }

    app.createContainer ()
    err := app.copyFileTree (
        strings.Join ([]string{ os.Getenv("GOPATH"), templatesPath }, string(os.PathSeparator)), app.name)
    if err != nil {
        log.Fatal (err)
    }
}


/**
 *  Creating root folder for new application
 */
func (app *Application) createContainer () {
    err := os.Mkdir(app.name, 0776); if err == nil {
        fmt.Println(green, "  create:", reset, app.name)
    } else {
        log.Fatal (err)
    }
}


func (app *Application) copyFileTree (from, to string) error {
    file, err := os.Open (from); if err != nil {
        return err
    }
    files, err := file.Readdir (0); if err != nil {
        return err
    }
    for _, f := range files {

        fromFilePath := strings.Join ([]string { from, f.Name () }, string(os.PathSeparator))
        toFilePath := strings.Join ([]string { to, f.Name () }, string(os.PathSeparator))

        if f.IsDir () == true {

            // copying folders
            err := os.Mkdir(toFilePath, f.Mode ()); if err != nil {
                return err
            } else {
                fmt.Println(green, "  create:", reset, toFilePath)
            }

            // scanning next level
            newTo := strings.Join ([]string { to, f.Name () }, string (os.PathSeparator))
            err = app.copyFileTree (fromFilePath, newTo); if err != nil {
                return err
            }
        } else {

            // copying files
            fileData, err := ioutil.ReadFile (fromFilePath); if err != nil {
                return err
            }
            err = ioutil.WriteFile (toFilePath, []byte(fileData), f.Mode ()); if err != nil {
                return err
            } else {
                fmt.Println(green, "  create:", reset, toFilePath)
            }
        }
    }

    return nil
}

// vim: noai:ts=4:sw=4
