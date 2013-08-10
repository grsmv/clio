package core

package (
    "strconv"
    "net/http"
    "github.com/pallada/clio/helpers"
    "github.com/daaku/go.grace/gracehttp"
    "fmt"
    "os"
)

func requestHandler () http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

        // setting up package variable to use outside the package
        ctx = context { ResponseWriter: w, Request: req }

        // setting up default headers
        setHeaders (w, req)

        router (w, req)
    })
    return mux
}


func Run (settings map[string]interface {}) {

    port := strconv.Itoa(settings["port"].(int))
    pidPath := settings["pid-file"].(string)

    // välkommen message
    fmt.Printf ("\n%sClio running. Port: %d, pid: %d%s \n%s\n\n",
        colours.green, settings["port"].(int), os.Getpid (), colours.reset,
        "For furter information please visit\nhttps://github.com/pallada/clio")

    // process-centric routines
    helpers.CreatePidFile (pidPath) // fix this!
    helpers.HandleSignals ()

    gracehttp.Serve (
        &http.Server { Addr: ":" + port, Handler: requestHandler () },
    )
}


// vim: noai:ts=4:sw=4
