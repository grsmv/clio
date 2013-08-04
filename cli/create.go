package cli

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

const templatesPath = "src/github.com/pallada/clio/templates"

type Application struct {
    name string
}

/**
 *  High level abstraction for creating new app
 *  (called from cli.Route())
 */
func Create (appName string) {
    app := Application {name: appName}

    app.createContainer ()
    app.createSubdirectories ()
    app.createConfigFiles ()
    app.createCoreJavascriptFiles ()
}

/**
 *  Creating root folder for new application
 */
func (app *Application) createContainer () {
    err := os.Mkdir(app.name, 0776)
    if err == nil {
        fmt.Println(green, "  create:", reset, app.name)
    } else {
        fmt.Println(red, "Terrible things happened:", reset)
        fmt.Println("   ", err)
        os.Exit(1)
    }
}

/**
 *  Creating all needed folders inside application's root folder
 */
func (app *Application) createSubdirectories () {
    folders := []string{
        "app",
        "app/assets",
        "app/assets/javascripts",
        "app/assets/stylesheets",
        "app/controllers",
        "app/helpers",
        "app/routes",
        "app/views",
        "bin",
        "config",
        "log",
        "public",
        "public/javascripts",
        "public/javascripts/vendor",
        "public/stylesheets",
        "tmp",
        "tmp/pids"}

    for i := range folders {
        path := strings.Join([]string{app.name, folders[i]}, string(os.PathSeparator))
        err := os.Mkdir(path, 0776)
        if err == nil {
            fmt.Println(green, "  create:", reset, path)
        }
    }
}

/**
 *  Creating needed config files
 */
func (app *Application) createConfigFiles () {
    app.createFilesFromTemplates ("config",
        []string{
            "application.go",
            "assets.yml",
            "dependencies",
            "procfile.yml"})
}

/**
 * Placing core JavaScript vendor files in application
 */
func (app *Application) createCoreJavascriptFiles () {
    coreComponents := []string{"underscore", "backbone"}
    for _, name := range coreComponents {
        app.createFilesFromTemplates ("public/javascripts/vendor/" + name, []string{ name + ".js"})
    }
}

/**
 *  Creating files copying from package templates
 *    @param <String> templatesSubfolder - directory containing listed templates
 *    @param <Array> files - list of templates to copy
 */
func (app *Application) createFilesFromTemplates (templatesSubfolder string, files []string) {
    sep := string(os.PathSeparator)

    for _, fileName := range files {
        parentFolder := strings.Join([]string{
            os.Getenv("GOPATH"),
            templatesPath,
            templatesSubfolder}, sep)

        parentFolderInApp := strings.Join([]string{
            app.name,
            templatesSubfolder}, sep)

        // checking if parent folder for file in app exists
        if _, err := os.Stat(parentFolderInApp); os.IsNotExist(err) {
            os.MkdirAll(parentFolderInApp, 0744)
        }

        // reading original template file and writing file into app
        filePath := parentFolder + sep + fileName
        fileData, err := ioutil.ReadFile(filePath)

        if err == nil {
            path := parentFolderInApp + sep + fileName
            ioutil.WriteFile(path, []byte(fileData), 0644)
            fmt.Println(green, "  create:", reset, path)
        } else {
            fmt.Println(red, "error:", reset, err)
        }
    }
}

// vim: noai:ts=4:sw=4
