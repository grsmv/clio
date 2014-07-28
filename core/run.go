package core

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/grsmv/clio/helpers"
)

var (
	AppSettings     = make(map[string]interface{})
	routes          = make(map[string]map[string]func() string)
	splat           = make(map[string]string)
	query           = make(map[string]string)
	contextInstance = context{}
	development     = false
)

type context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	ResponseCode   int
}

/**
 *  Creating namespace to place routes by specific http method
 */
func init() {
	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	for index := range methods {
		routes[methods[index]] = make(map[string]func() string)
	}

	// geting data from CLI
	flag.BoolVar(&development, "development", false, "")
	flag.Parse()
}

func Handler(w http.ResponseWriter, req *http.Request) {
	// setting up package variable to use outside the package
	contextInstance = context{ResponseWriter: w, Request: req}

	// setting up default headers
	setHeaders(w, req)

	Router(w, req)
}

func Run(settings map[string]interface{}) {

	// making application's settings accessible to whole package
	AppSettings = settings

	// basic assets management
	if settings["manage-assets"].(bool) {
		fs := http.FileServer(http.Dir("static"))
		http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	}

	http.HandleFunc("/", Handler)

	// initializing all registered websockets
	InitializeWebsockets()

	port := strconv.Itoa(settings["port"].(int))

	message := "Clio server started at " + strconv.Itoa(settings["port"].(int)) + " port"

	if development {
		message += " in development mode"
	}

	log.Println(message)
	http.ListenAndServe(":"+port, nil)
}

/**
 *  Finding correct handler to certain method:path
 */
func Router(w http.ResponseWriter, req *http.Request) {

	// splitting whole path into parts
	path, queryString := helpers.SplitPath(req.URL.String())
	routeFound := false

	// finding correct handler
	for rawPattern, _ := range routes[req.Method] {
		pattern := helpers.PreparePattern(rawPattern)

		if pattern.MatchString(path) {
			routeFound = true

			// homage to Sinatra's splat
			splat = helpers.ParseSplat(pattern, path)

			// filling query
			query = helpers.ParseQuery(queryString)

			// prepearing buffer
			buffer := new(bytes.Buffer)

			// default values for 'Before' or 'After' hooking
			hooksAvailable := false
			contextInstance.ResponseCode = 200

			// calling before action
			if BeforeActionStore[req.Method][rawPattern] != nil {
				hooksAvailable = true
				buffer.Write([]byte(BeforeActionStore[req.Method][rawPattern]()))
			}

			// calling matched handler
			if !hooksAvailable || Context().ResponseCode == 200 {
				buffer.Write([]byte(routes[req.Method][rawPattern]()))
			}

			// flushing buffer to ResponseWriter
			buffer.WriteTo(w)

			// terminal debugging
			if Verbose() {
				log.Printf("%s %s\n", req.Method, req.URL.String())
			}
			break
		}
	}

	if !routeFound {
		NotFound(w, req)
	}
}

func Verbose() bool {
	return AppSettings["verbose-output"] != nil && AppSettings["verbose-output"].(bool) == true
}
