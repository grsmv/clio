package cli

import (
    "fmt"
    "github.com/cliohq/clio/vendor/fsnotify"
    "log"
    "net/rpc"
    "regexp"
    "sync"
)

func Watch () {
    watcher, err := fsnotify.NewWatcher (); if err != nil {
        log.Fatal (err)
    }

    var wg sync.WaitGroup

    wg.Add (1)

    go func () {
        for {
            select {
            case ev := <-watcher.Event:
                match, _ := regexp.Match("\\.(go|template)$", []byte(ev.Name))

                if match && ev.IsModify () {
                    fmt.Print (ev.Name + ". app rebuild...")

                    // send signal to app relaunch
                    RelaunchProcessCall ()
                }
            }
        }
        wg.Done ()
    } ()

    // todo: move list of directories to watch to settings
    watcher.Watch ("./app/controllers")
    watcher.Watch ("./app/helpers")
    watcher.Watch ("./app/views")
    watcher.Watch ("./app/routes")
    watcher.Watch ("./config")

    wg.Wait()
}


func RelaunchProcessCall () {
    client, err := rpc.DialHTTP ("tcp", "localhost:31000"); if err != nil {
        log.Fatal ("dialing error", err)
    }

    err = client.Call ("Server.RelaunchProcess", nil, nil); if err != nil {
        log.Fatal ("server error:", err)
    }
}

// vim: noai:ts=4:sw=4
