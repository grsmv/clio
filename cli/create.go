package cli

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
)


type Application struct {
    name string
}

var (
    GOPATH = os.Getenv ("GOPATH")
    slash = string (os.PathSeparator)
    GOPATH_SRC = GOPATH + slash + "src" + slash
)


/**
 *  High level abstraction for creating new app
 *  (called from cli.Route())
 */
func Create (appName string) {

    checkGopath ()
    checkContainer (appName)

    app := Application { name: appName }

    app.createContainer ()

    err := app.copyFileTree (
        // source
        strings.Join (
            []string { GOPATH, applicationTemplatesPath },
            slash),

        // destination
        GOPATH_SRC + app.name)

    if err != nil {
        log.Fatal (err)
    }
}


/**
 * Cheking if GOPATH setted properly
 */
func checkGopath () {
    if len (GOPATH) == 0 {
        log.Fatal ("GOPATH is empty. Please fix this")
    }
}


/**
 * Checing if folder with desirable name already exists
 */
func checkContainer (appName string) {
    destination := GOPATH_SRC + appName
    _, err := os.Stat (destination); if os.IsExist (err) {
        log.Fatal ("folder " + destination + " already exists")
    }
}


/**
 *  Creating root folder for new application
 */
func (app *Application) createContainer () {
    appPath := GOPATH_SRC + app.name
    err := os.Mkdir(appPath, 0776); if err == nil {
        fmt.Println(green, "  create:", reset, appPath)
    } else {
        log.Fatal (err)
    }
}


/**
 * Creating an application sceleton from templates/application
 */
func (app *Application) copyFileTree (from, to string) error {
    file, err := os.Open (from); if err != nil {
        return err
    }
    files, err := file.Readdir (0); if err != nil {
        return err
    }
    for _, f := range files {

        fromFilePath := strings.Join ([]string { from, f.Name () }, slash)
        toFilePath := strings.Join ([]string { to, f.Name () }, slash)

        if f.IsDir () == true {

            // copying folders
            err := os.Mkdir(toFilePath, f.Mode ()); if err != nil {
                return err
            } else {
                fmt.Println(green, "  create:", reset, toFilePath)
            }

            // scanning next level
            newTo := strings.Join ([]string { to, f.Name () }, slash)
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
