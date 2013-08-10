package core

import (
    "strconv"
    "net/http"
    "github.com/pallada/clio/helpers"
    "github.com/pallada/clio/vendor/gracehttp"
)

func requestHandler () http.Handler {
    mux := http.NewServeMux()
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
    pidPath := settings["pid-file"].(string)

    // process-centric routines
    helpers.CreatePidFile (pidPath) // fix this!

    gracehttp.Serve (
        &http.Server { Addr: ":" + port, Handler: requestHandler () },
    )
}


// vim: noai:ts=4:sw=4
