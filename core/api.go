package core

import (
    "strconv"
    "net/http"
)

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

// ------------------- app runner -------------------

func Run (port int) {
    http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        // setting up package variable to use outside the package
        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        router(w, req)
    })

    http.ListenAndServe(":" + strconv.Itoa(port), nil)
}


// vim: noai:ts=4:sw=4
