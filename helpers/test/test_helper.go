package test

import (
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
)

type Response struct {
    Server        *httptest.Server
    Body          string
    StatusCode    int
    Header        http.Header
    ContentLength int64
}


func NewResponse (server *httptest.Server) Response {
    return Response { Server: server }
}


func (th *Response) Get (url string, cookies []http.Cookie, headers map[string][]string) Response {
    return genericRequest ("GET", url, cookies, headers)
}


func (th *Response) Post (url string, cookies []http.Cookie, headers map[string][]string) Response {
    return genericRequest ("POST", url, cookies, headers)
}


func (th *Response) Put (url string, cookies []http.Cookie, headers map[string][]string) Response {
    return genericRequest ("PUT", url, cookies, headers)
}


func (th *Response) Delete (url string, cookies []http.Cookie, headers map[string][]string) Response {
    return genericRequest ("DELETE", url, cookies, headers)
}


// todo: set headers
func genericRequest (method, url string, cookies []http.Cookie, headers map[string][]string) Response {
    request, _ := http.NewRequest(
        method,
        url,
        io.MultiReader())

    // note: http://golang.org/pkg/net/http/#Request.AddCookie
    if cookies != nil {
        for _, cookie := range cookies {
            request.AddCookie(&cookie)
        }
    }

    if headers != nil {
        request.Header = headers
    }

    client := http.Client {}
    response, _ := client.Do (request)
    body, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    return Response {
        StatusCode:     response.StatusCode,
        Body:           string(body),
        Header:         response.Header,
        ContentLength:  response.ContentLength,
    }
}
