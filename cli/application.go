package cli

import (
    "github.com/cliohq/clio/helpers"
    "log"
    "net"
    "net/http"
    "net/rpc"
    "os/exec"
    "strconv"
)

const applicationProcessCode = "app"

var LaunchedProcesses map[string]*exec.Cmd

type Server int

type Args struct {
    ProcName string
}

type ProcessInfo struct {
    Path string
    Args []string
    Pid int
}


func init () {
    LaunchedProcesses = make(map[string]*exec.Cmd)
}


func LaunchTcpServer () {
    server := new (Server)
    rpc.Register (server)
    rpc.HandleHTTP ()
    l, err := net.Listen ("tcp", ":" + strconv.Itoa(tcpIPCPort)); if err != nil {
        log.Fatal ("listen error:", err)
    }
    http.Serve (l, nil)
}


func (t *Server) RelaunchProcess (args *Args, reply *int) error {
    appProc := LaunchedProcesses[applicationProcessCode]

    // backupig application's process info
    processBackup := BackupProcess (appProc)

    // killing old app's process
    err := appProc.Process.Kill(); if err != nil {
        log.Fatal ("cli.RelaunchProcess():", err)  /////////// debug
    }

    // rebuilding application
    helpers.ApplicationRebuild () // sync

    // Relaunch application
    newApplicationProc := exec.Command (processBackup.Path, processBackup.Args...)

    // streaming output from new app instance
    stdOut, _ := newApplicationProc.StdoutPipe()
    stdErr, _ := newApplicationProc.StderrPipe()

    StreamOutput (stdOut, stdErr, applicationProcessCode)

    go newApplicationProc.Start ()

    // updating `LaunchedProcesses`
    LaunchedProcesses[applicationProcessCode] = newApplicationProc

    return nil
}


func BackupProcess (command *exec.Cmd) ProcessInfo {
    return ProcessInfo { command.Path, command.Args, command.Process.Pid }
}


// vim: noai:ts=4:sw=4
