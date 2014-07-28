package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	contentTypes = map[string]string{
		"css":  "text/css",
		"csv":  "text/csv",
		"gif":  "image/gif",
		"jpeg": "image/jpeg",
		"jpg":  "image/jpeg",
		"js":   "application/javascript",
		"json": "application/json",
		"png":  "image/png",
		"svg":  "image/svg+xml",
		"txt":  "text/plain",
		"xml":  "text/xml"}

	defaultHeaders = map[string]string{
		"Content-Type":   "text/html",
		"Accept-Charset": "utf-8"}
)

func Redirect(url string) {
	http.Redirect(Context().ResponseWriter, Context().Request, url, http.StatusFound)
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template404, err := ioutil.ReadFile("public/404.html")
	if err != nil {
		template404 = []byte("404 Not Found")
	}
	w.WriteHeader(404)
	fmt.Fprintln(w, string(template404))
}

func setHeaders(w http.ResponseWriter, req *http.Request) {

	//FIXME: copying map on each request is bad
	var requestHeaders = make(map[string]string)
	for k, v := range defaultHeaders {
		requestHeaders[k] = v
	}

	// getting request format
	pattern, _ := regexp.Compile("\\.(css|csv|gif|jpeg|jpg|js|json|png|svg|txt|xml)$")
	if pattern.MatchString(req.URL.String()) {
		requestType := pattern.FindAllStringSubmatch(req.URL.String(), 1)[0][1]
		requestHeaders["Content-Type"] = contentTypes[requestType]
	}

	// setting up custom headers
	for key, value := range requestHeaders {
		w.Header().Set(key, value)
	}
}
