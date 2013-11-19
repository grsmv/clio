package core

import (
    "fmt"
    "github.com/cliohq/clio/helpers"
    "net/http"
    "log"
)

var (
    AppSettings = make(map[string]interface{})
    routes      = make(map[string] map[string] func () string)
    splat       = []string{}
    params      = make(map[string]string)
    ctx         = context {}
)

type context struct {
    Request *http.Request
    ResponseWriter http.ResponseWriter
}


/**
 *  Creating namespace to place routes by specific http method
 */
func init () {
    methods := []string{"GET", "POST", "PUT", "DELETE", "WS"}
    for index := range methods {
        routes[methods[index]] = make (map[string] func () string)
    }
}


/**
 *  Finding correct handler to certain method:path
 */
func Router (w http.ResponseWriter, req *http.Request) {

    // splitting whole path into parts
    path, paramsString := helpers.SplitPath(req.URL.String())

    routeFound := false

    method := req.Method
    if IsWebsocket() {
        method = "WS"
    }

    // finding correct handler
    for rawPattern, _ := range routes[method] {
        pattern := helpers.PreparePattern(rawPattern)

        if pattern.MatchString(path) {
            routeFound = true

            // homage to Sinatra's splat
            splat = pattern.FindAllStringSubmatch(path, 100)[0][1:]

            // filling params
            params = helpers.ParseParams(paramsString)

            // calling matched handler
            fmt.Fprintln(w, routes[method][rawPattern]())

            // terminal debugging
            if AppSettings["verbose-output"] != nil && AppSettings["verbose-output"].(bool) == true {
                log.Printf ("%s %s\n", method, req.URL.String())
            }
            break
        }
    }
    if !routeFound {
        NotFound(w, req)
    }
}

// vim: noai:ts=4:sw=4
