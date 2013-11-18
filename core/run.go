package core

import (
    "net/http"
    "strconv"
    "log"
)

func requestHandler (settings map[string]interface{}) {

    // basic assets management
    if settings["manage-assets"].(bool) {
        fs := http.FileServer(http.Dir("public"))
        http.Handle("/assets/", http.StripPrefix("/assets/", fs))
    }

    http.HandleFunc("/", Handler)
}


func Handler (w http.ResponseWriter, req *http.Request) {
    // setting up package variable to use outside the package
    ctx = context { ResponseWriter: w, Request: req }

    // setting up default headers
    setHeaders (w, req)

    Router (w, req)
}


func Run (settings map[string]interface {}) {

    // making application's settings accessible to whole package
    AppSettings = settings

    requestHandler (settings)

    port := strconv.Itoa(settings["port"].(int))

    log.Println ("Clio server started at", AppSettings["port"].(int), "port")
    http.ListenAndServe (":" + port, nil)
}


// todo: share settings by IPC

// vim: noai:ts=4:sw=4
