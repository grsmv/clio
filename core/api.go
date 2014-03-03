package core

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
)

func init() {
	BeforeActionStore = map[string]map[string]func() string{
		"GET":     map[string]func() string{},
		"POST":    map[string]func() string{},
		"PUT":     map[string]func() string{},
		"DELETE":  map[string]func() string{},
		"OPTIONS": map[string]func() string{},
	}
}

// helpers for 'before actions' implementation ------------

var BeforeActionStore map[string]map[string]func() string

type ActionHandler struct {
	Method  string
	Pattern string
}

func (ah ActionHandler) Before(do func() string) {
	BeforeActionStore[ah.Method][ah.Pattern] = do
	// return ah
}

// --------------------------------------------------------

// core utilities -----------------------------------------

func Context() context {
	return ctx
}

func Splat() map[string]string {
	return splat
}

func Query() map[string]string {
	return query
}

func Params(q string) string {
	return ctx.Request.FormValue(q)
}

func SetHeader(key, value string) {
	Context().ResponseWriter.Header().Set(key, value)
}

// --------------------------------------------------------

// route-declaring methods --------------------------------

func Get(pattern string, handler func() string) ActionHandler {
	ctx.ResponseCode = 200
	routes["GET"][pattern] = handler
	return ActionHandler{"GET", pattern}
}

func Post(pattern string, handler func() string) ActionHandler {
	ctx.ResponseCode = 200
	routes["POST"][pattern] = handler
	return ActionHandler{"POST", pattern}
}

func Put(pattern string, handler func() string) ActionHandler {
	ctx.ResponseCode = 200
	routes["PUT"][pattern] = handler
	return ActionHandler{"PUT", pattern}
}

func Delete(pattern string, handler func() string) ActionHandler {
	ctx.ResponseCode = 200
	routes["DELETE"][pattern] = handler
	return ActionHandler{"DELETE", pattern}
}

func Options(pattern string, handler func() string) ActionHandler {
	ctx.ResponseCode = 200
	routes["OPTIONS"][pattern] = handler
	return ActionHandler{"OPTIONS", pattern}
}

// --------------------------------------------------------

// HTTP error codes shortcuts -----------------------------

func AccessDenied() {
	ctx.ResponseCode = 403
	SetHeader("WWW-Authenticate", "Basic Realm=\"My Realm\"")
	http.Error(Context().ResponseWriter, "Authentication required", 403)
}

func BadRequest() {
	ctx.ResponseCode = 400
	http.Error(Context().ResponseWriter, "Bad Request", 400)
}

func NotFoundError() {
	ctx.ResponseCode = 404
	http.Error(Context().ResponseWriter, "Not Found", 404)
}

func Conflict() {
	ctx.ResponseCode = 409
	http.Error(Context().ResponseWriter, "Conflict", 409)
}

// --------------------------------------------------------

// other helpers ------------------------------------------

func RequestBody() string {
	body, _ := ioutil.ReadAll(Context().Request.Body)
	return string(body)
}

// Populating given empty instance of certain class with
// form data
// Example usage:
//     Populate (new(User))
func Populate(instance interface{}) interface{} {
	decoder := schema.NewDecoder()
	ctx.Request.ParseForm()
	decoder.Decode(instance, ctx.Request.Form)
	return instance
}

// --------------------------------------------------------

// shortcut for json returning ----------------------------

func Json(obj interface{}) string {
	SetHeader("Content-Type", "application/json")

	b, _ := json.Marshal(obj)
	return string(b)
}

// --------------------------------------------------------
