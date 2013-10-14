package helpers

import (
    "os/exec"
    "log"
)


func ApplicationRebuild () {
    goBinPath, err := exec.LookPath ("go"); if err != nil {
        log.Fatal ("helpers.ApplicationRebuild():", err) ////// debug
    }

    command := exec.Command (goBinPath, "build", "application.go")

    err = command.Start(); if err != nil {
        log.Fatal ("helpers.ApplicationRebuild():", err) ////// debug
    }

    command.Wait ()
}

// vim: noai:ts=4:sw=4
