package main

import (
    "code.google.com/p/go.net/websocket"
    "flag"
    "log"
    "net/http"
    "text/template"
)

type hub struct {
    // registered connections
    connections map[connection]bool

    // inbound messages from the connections
    broadcast chan string

    // register requests from the connections
    register chan connection

    // unregister requests from connections
    unregister chan connection
}

var h = hub {
    broadcast:   make (chan string),
    register:    make (chan connection),
    unregister:  make (chan connection),
    connections: make (map[connection]bool),
}

func (h *hub) run () {
    for {
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            delete (h.connections, c)
            close (c.send)
        case m := <-h.broadcast:
            for c := range h.connections {
                select {
                case c.send <- m:
                default:
                    delete (h.connections, c)
                    close (c.send)
                    go c.ws.Close ()
                }
            }
        }
    }
}

type connection struct {
    ws *websocket.Conn
    send chan string
}

func (c *connection) reader () {
    for {
        var message [20]byte
        _, err := c.ws.Read(message[:])
        if err != nil { break }
        h.broadcast <- string(message[:])
    }
    c.ws.Close()
}

func (c *connection) writer () {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil { break }
    }
    c.ws.Close ()
}

func wsHandler (ws *websocket.Conn) {
    c := connection { send: make (chan string, 256), ws: ws }
    h.register <- c
    defer func () { h.unregister <- c } ()
    go c.writer ()
    c.reader ()
}

func Socket (path string) {
    http.Handle (path, websocket.Handler (wsHandler))
}

var addr = flag.String ("addr", ":8080", "http service address")
var homeTempl = template.Must (template.ParseFiles ("index.html"))

func homeHandler (c http.ResponseWriter, req *http.Request) {
    homeTempl.Execute (c, req.Host)
}

func main () {
    flag.Parse ()
    go h.run ()
    http.HandleFunc ("/", homeHandler)
    // http.Handle ("/ws", websocket.Handler (wsHandler))
    // http.Handle ("/ws2", websocket.Handler (wsHandler))

    Socket ("/ws")
    Socket ("/ws2")

    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal ("ListenAndServe:", err)
    }
}

// vim: noai:ts=4:sw=4
