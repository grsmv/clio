package core

import (
    "strconv"
    "net/http"
    "github.com/pallada/clio/helpers"
    "encoding/json"
    "fmt"
)

// ------------------- before everything -------------------

var colours Colours

func init () {
    colours = Colours {}
    colours.init ()
}

// ------------------- utilities -------------------

func Context () context {
    return ctx
}


func Splat () []string {
    return splat
}


func Params () map[string]string {
    return params
}


func SetHeader (key, value string) {
    Context().ResponseWriter.Header().Set(key, value)
}

// ------------------- rest methods -------------------

func Get (pattern string, handler func () string) {
    routes["GET"][pattern] = handler;
}


func Post (pattern string, handler func () string) {
    routes["POST"][pattern] = handler;
}


func Put (pattern string, handler func () string) {
    routes["PUT"][pattern] = handler;
}


func Delete (pattern string, handler func () string) {
    routes["DELETE"][pattern] = handler;
}

// ------------------- json helper -------------------

func Json (obj interface{}) string {
    SetHeader("Content-Type", "application/json")

    b, _ := json.Marshal(obj)
    return string(b)
}

// ------------------- app runner -------------------

func Run (settings map[string]interface {}) {

    // välkommen message
    fmt.Printf ("\n%sClio running. Port %d%s\n%s\n\n",
        colours.green, settings["port"].(int), colours.reset,
        "For furter information please visit\nhttps://github.com/pallada/clio")

    http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        // setting up package variable to use outside the package
        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        router (w, req)
    })

    // process-centric routines
    helpers.CreatePidFile (settings["pid-file"].(string))
    helpers.HandleSignals ()

    // running and handling income
    http.ListenAndServe(":" + strconv.Itoa(settings["port"].(int)), nil)
}


// vim: noai:ts=4:sw=4
