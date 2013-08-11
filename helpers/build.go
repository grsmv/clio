package helpers

import (
    "strings"
    "io/ioutil"
    "os/exec"
    "log"
)

func Build () {
    Exec ("go build application.go")
}


func Euthanasia () {
    data, err := ioutil.ReadFile ("tmp/pids/clio.pid")
    if err != nil {
        log.Fatal (err)
    }
    Exec ("kill -USR2 " + string(data))
}


func Exec (call string) {
    callParts := strings.Split(call, " ")
    command := exec.Command(callParts[0], callParts[1:]...)
    err := command.Start()
    if err != nil {
        log.Fatal (err)
    }
}

// vim: noai:ts=4:sw=4
