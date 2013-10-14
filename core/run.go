package core

import (
    "net/http"
    "strconv"
    "fmt"
)

func requestHandler (settings map[string]interface{}) {

    serveStatic := settings["manage-assets"].(bool)
    if serveStatic {
        fs := http.FileServer(http.Dir("public"))
        http.Handle("/assets/", http.StripPrefix("/assets/", fs))
    }

    http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        // setting up package variable to use outside the package
        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        router (w, req)
    })
}


func Run (settings map[string]interface {}) {
    requestHandler (settings)

    port := strconv.Itoa(settings["port"].(int))

    fmt.Println ("Clio server started at", settings["port"].(int), "port")
    http.ListenAndServe (":" + port, nil)
}


// todo: share settings by IPC

// vim: noai:ts=4:sw=4
