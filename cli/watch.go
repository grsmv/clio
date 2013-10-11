package cli

import (
    "github.com/cliohq/clio/vendor/fsnotify"
    /* "github.com/cliohq/clio/cli" */
    "log"
    "fmt"
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
                    fmt.Println (ev.Name + " changed. restarting an app")

                    Relaunch ("")

                    // helpers.Build ()
                    // helpers.Euthanasia ()
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

// vim: noai:ts=4:sw=4
