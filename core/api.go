package core

import (
    "encoding/json"
    "github.com/gorilla/schema"
    "net/http"
    "io/ioutil"
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


func Splat () map[string]string {
    return splat
}


func Query () map[string]string {
    return query
}

func Params (q string) (string) {
    return ctx.Request.FormValue(q)
}


func SetHeader (key, value string) {
    Context().ResponseWriter.Header().Set(key, value)
}

func AccessDenied () {
    SetHeader("WWW-Authenticate", "Basic Realm=\"My Realm\"")
    http.Error(Context().ResponseWriter, "Authentication required", 403)
}

func RequestBody () string {
    body, _ := ioutil.ReadAll(Context().Request.Body)
    return string(body)
}


// Populating given empty instance of certain class with
// form data
// Example usage:
//     Populate (new(User))
func Populate (instance interface{}) interface{} {
    decoder := schema.NewDecoder()
    ctx.Request.ParseForm()
    decoder.Decode(instance, ctx.Request.Form)
    return instance
}

// ------------------- rest methods -------------------

func Get (pattern string, handler func () string) {
    routes["GET"][pattern] = handler
}


func Post (pattern string, handler func () string) {
    routes["POST"][pattern] = handler
}


func Put (pattern string, handler func () string) {
    routes["PUT"][pattern] = handler
}


func Delete (pattern string, handler func () string) {
    routes["DELETE"][pattern] = handler
}

func Options (pattern string, handler func () string) {
    routes["OPTIONS"][pattern] = handler
}

// ------------------- json helper -------------------

func Json (obj interface{}) string {
    SetHeader("Content-Type", "application/json")

    b, _ := json.Marshal(obj)
    return string(b)
}


// vim: noai:ts=4:sw=4
