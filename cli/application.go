package cli

import (
    "os/exec"
    "net/rpc"
    "net/http"
    "net"
    "log"
    "github.com/cliohq/clio/helpers"
)

type Server int

type Args struct {
    ProcName string
}

type ProcessInfo struct {
    Path string
    Args []string
    Pid int
}


var LaunchedProcesses map[string]*exec.Cmd


func init () {
    LaunchedProcesses = make(map[string]*exec.Cmd)
}


func LaunchTcpServer () {
    server := new (Server)
    rpc.Register (server)
    rpc.HandleHTTP ()
    l, err := net.Listen ("tcp", ":31000"); if err != nil {
        log.Fatal ("listen error:", err)
    }
    http.Serve (l, nil)
}


func (t *Server) RelaunchProcess (args *Args, reply *int) error {
    appProc := LaunchedProcesses["app"]

    // backupig application's process info
    processBackup := BackupProcess (appProc)

    // killing old app's process
    err := appProc.Process.Kill(); if err != nil {
        log.Fatal ("RelaunchProcess: ", err)  /////////// debug
    }

    // rebuilding application
    helpers.ApplicationRebuild () // sync

    // Relaunch application
    newApplicationProc := exec.Command (processBackup.Path, processBackup.Args...)

    go newApplicationProc.Start ()

    // updating `LaunchedProcesses`
    LaunchedProcesses["app"] = newApplicationProc

    return nil
}


func BackupProcess (command *exec.Cmd) ProcessInfo {
    return ProcessInfo { command.Path, command.Args, command.Process.Pid }
}


// vim: noai:ts=4:sw=4
