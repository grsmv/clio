package clio

import (
    "net/http"
    "fmt"
    "strconv"
)

var (
    routes = make(map[string] map[string] func () string)
    splat = []string{}
    params = make(map[string]string)
    ctx = context {}
)

type context struct {
    Request *http.Request
    ResponseWriter http.ResponseWriter
}


func init () {
    // creating namespace to place routes by specific http method
    methods := []string{"GET", "POST", "PUT", "DELETE"}
    for index := range methods {
        routes[methods[index]] = make (map[string] func () string)
    }
}


func Context () context {
    return ctx
}


func Splat () []string {
    return splat
}


func Params () map[string]string {
    return params
}


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


func Run (port int) {

    http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        // finding correct handler
        for rawPattern, _ := range routes[req.Method] {
            pattern := prepearePattern(rawPattern)

            // splitting whole path into parts
            path, paramsString := splitPath(req.URL.String())

            if pattern.MatchString(path) {

                // homage to Sinatra's splat
                splat = pattern.FindAllStringSubmatch(path, 100)[0][1:]

                // filling params
                params = parseParams(paramsString)

                // calling matched handler
                fmt.Fprintln(w, routes[req.Method][rawPattern]())

                // terminal debugging
                fmt.Println(req.Method, req.URL.String())
                break
            }
        }
    })

    http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

