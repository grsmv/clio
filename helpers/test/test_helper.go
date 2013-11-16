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


func NewResponse(server *httptest.Server) Response {
    return Response { Server: server }
}


func (r *Response) Get(url string, cookies []http.Cookie, headers map[string][]string) Response {
    return r.genericRequest ("GET", url, cookies, headers)
}


func (r *Response) Post(url string, cookies []http.Cookie, headers map[string][]string) Response {
    return r.genericRequest ("POST", url, cookies, headers)
}


func (r *Response) Put(url string, cookies []http.Cookie, headers map[string][]string) Response {
    return r.genericRequest ("PUT", url, cookies, headers)
}


func (r *Response) Delete(url string, cookies []http.Cookie, headers map[string][]string) Response {
    return r.genericRequest ("DELETE", url, cookies, headers)
}


func (r *Response) genericRequest(method, url string, cookies []http.Cookie, headers map[string][]string) Response {
    request, _ := http.NewRequest(
        method,
        r.Server.URL  + url,
        io.MultiReader())

    // note: http://golang.org/pkg/net/http/#Request.AddCookie
    // todo: test with custom cookies
    if cookies != nil {
        for _, cookie := range cookies {
            request.AddCookie(&cookie)
        }
    }

    // todo: test with custom headers
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
