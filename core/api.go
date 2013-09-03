package core

import (
    "encoding/json"
)

// ------------------- initializing colour marks -------------------

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


// vim: noai:ts=4:sw=4
