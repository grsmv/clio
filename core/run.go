package core

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "strconv"
    "log"
)

func httpHandler (w http.ResponseWriter, req *http.Request) {
    // setting up package variable to use outside the package
    ctx = context { ResponseWriter: w, Request: req }

    if IsWebsocket() {
        websocket.Handler(func(ws *websocket.Conn) {
          Handler(w, req, ws)
        }).ServeHTTP(w, req)
    } else {
        Handler(w, req, nil)
    }
}

func Handler (w http.ResponseWriter, req *http.Request, ws *websocket.Conn) {

    //req.Websocket = ws

    // setting up default headers
    setHeaders (w, req)

    Router (w, req)
}

func requestHandler (settings map[string]interface{}) {

    // basic assets management
    if settings["manage-assets"].(bool) {
        fs := http.FileServer(http.Dir("public"))
        http.Handle("/assets/", http.StripPrefix("/assets/", fs))
    }

    http.HandleFunc("/", httpHandler)

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
