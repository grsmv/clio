package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Application struct {
	Name string
}

/**
 *  High level abstraction for creating new app
 *  (called from cli.Route())
 */
func Create(appName string) {

	checkGopath()
	checkContainer(appName)

	app := Application{Name: appName}

	app.createContainer()

	err := app.copyFileTree(
		GOPATH+slash+applicationTemplatesPath,
		GOPATH_SRC+app.Name,
	)

	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Cheking if GOPATH setted properly
 */
func checkGopath() {
	if len(GOPATH) == 0 {
		log.Fatal("GOPATH is empty. Please fix this")
	}
}

/**
 * Checing if folder with desirable name already exists
 */
func checkContainer(appName string) {
	destination := GOPATH_SRC + appName
	_, err := os.Stat(destination)
	if os.IsExist(err) {
		log.Fatal("folder " + destination + " already exists")
	}
}

/**
 *  Creating root folder for new application
 */
func (app *Application) createContainer() {
	appPath := GOPATH_SRC + app.Name
	if err := os.Mkdir(appPath, 0776); err == nil {
		fmt.Println(green, "  create:", reset, appPath)
	} else {
		log.Fatal(err)
	}
}

/**
 *  Copying certain directory from applciation's templates
 *  to a new application's skeleton
 */
func (app *Application) copyDir(file os.FileInfo, destination, fromFilePath, toFilePath string) error {
	if err := os.Mkdir(toFilePath, file.Mode()); err == nil {
		fmt.Println(green, "  create:", reset, toFilePath)
	} else {
		return err
	}

	// scanning next level
	newTo := strings.Join([]string{destination, file.Name()}, slash)
	if err := app.copyFileTree(fromFilePath, newTo); err != nil {
		return err
	}

	return nil
}

/**
 *  Copying certain file from applciation's templates
 *  to a new application's skeleton
 */
func (app *Application) copyFile(file os.FileInfo, source, destination string) error {

	fileData, _ := ioutil.ReadFile(source)

	// detecting template and processing it
	if strings.HasSuffix(source, ".go.template") {
		var buffer bytes.Buffer
		tmpl := template.Must(template.New("app").Parse(string(fileData)))
		tmpl.Execute(&buffer, app)

		// overwriting file contents with data from processed template
		fileData = buffer.Bytes()

		// removing tmp suffixes from file names
		destination = strings.Replace(destination, ".template", "", -1)
	}

	if err := ioutil.WriteFile(destination, fileData, file.Mode()); err == nil {
		fmt.Println(green, "  create:", reset, destination)
	} else {
		return err
	}
	return nil
}

/**
 * Creating an application sceleton from templates/application
 */
func (app *Application) copyFileTree(from, to string) error {
	file, err := os.Open(from)
	if err != nil {
		return err
	}
	files, err := file.Readdir(0)
	if err != nil {
		return err
	}
	for _, f := range files {

		fromFilePath := strings.Join([]string{from, f.Name()}, slash)
		toFilePath := strings.Join([]string{to, f.Name()}, slash)

		if f.IsDir() == true {
			app.copyDir(f, to, fromFilePath, toFilePath)
		} else {
			app.copyFile(f, fromFilePath, toFilePath)
		}
	}

	return nil
}

// vim: noai:ts=4:sw=4
