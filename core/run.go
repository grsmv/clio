package core

import (
    "github.com/cliohq/clio/vendor/go.grace/gracehttp"
    "net/http"
    "strconv"
)

func requestHandler (settings map[string]interface{}) http.Handler {

    mux := http.NewServeMux()

    serveStatic := settings["manage-assets"].(bool)
    if serveStatic {
        fs := http.FileServer(http.Dir("public"))
        mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
    }

    mux.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        // setting up package variable to use outside the package
        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        router (w, req)
    })
    return mux
}


func Run (settings map[string]interface {}) {
    port := strconv.Itoa(settings["port"].(int))

    gracehttp.Serve (
        &http.Server { Addr: ":" + port, Handler: requestHandler (settings) },
    )
}


// vim: noai:ts=4:sw=4
