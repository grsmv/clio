package core

import (
    "strconv"
    "net/http"
    "github.com/pallada/clio/vendor/go.grace/gracehttp"
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

    gracehttp.Serve (
        &http.Server { Addr: ":" + port, Handler: requestHandler () },
    )
}


// vim: noai:ts=4:sw=4
