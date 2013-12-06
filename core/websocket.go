package core

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "log"
)

type ClientConnection struct {
    websocket *websocket.Conn
    clientIP string
}


var (
    // placeholer for all open connections, sorted by channel's name
    ActiveClients = make(map[string]map[ClientConnection]int)

    // list of all available channels
    Channels = []string{}
)


/**
 *  Registering new channel
 */
func Ws (channel string) {
    Channels = append(Channels, channel)
}


/**
 *  Initializing all registered websockets
 */
func InitializeWebsockets () {
    for _, channel := range Channels {
        ChannelInit(channel)
    }
}


/**
 *  Initializing new channel
 */
func ChannelInit (channel string) {

    ActiveClients[channel] = make(map[ClientConnection]int)

    http.Handle(channel, websocket.Handler(func(ws *websocket.Conn) {
        var err error
        var clientMessage string

        // cleanup on a server side
        defer func() {
            if err = ws.Close(); err != nil && Verbose () {
                log.Println("Websocket could not be closed", err.Error())
            }
        }()

        client := ws.Request().RemoteAddr

        if Verbose() {
            log.Print("Client connected:", client)
        }

        socketClient := ClientConnection { ws, client }
        ActiveClients[channel][socketClient] = 0

        if Verbose() {
            log.Println("Number of clients connected ...", len(ActiveClients[channel]))
        }

        // for loop so the websocket stays open otherwise
        // it will close after one Receive and Send
        for {

            // if we cannot Read - then connection is closed
            if err = websocket.Message.Receive(ws, &clientMessage); err != nil {

                if Verbose() {
                    log.Println("Websocket disconnected waiting", err.Error())
                }

                // removing the ws client connection from the list of our active clients
                delete(ActiveClients[channel], socketClient)

                if Verbose() {
                    log.Println("Number of clients still connected ...", len(ActiveClients[channel]))
                }
                return
            }

            // clientMessage = socketClient.clientIP + " Said: " + clientMessage // can be useful during debug
            for cs, _ := range ActiveClients[channel] {
                if err = websocket.Message.Send(cs.websocket, clientMessage); err != nil {

                    // we could not send message to a peer
                    if Verbose () {
                        log.Println("Could not send message to ", cs.clientIP, err.Error())
                    }
                }
            }
        }
    }))
}


/**
 * Sending message to all clients of a certain channel
 */
func WsSend (channel, message string) (err error) {
    for cs, _ := range ActiveClients[channel] {
        err = websocket.Message.Send(cs.websocket, message); if err != nil {
            log.Println(channel, "Could not send message to ", cs.clientIP, err.Error())
        }
    }
    return
}
