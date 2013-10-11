package cli

import (
    "os/exec"
    "net/rpc"
    "net/http"
    "net"
    "log"
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
    // procBackup := BackupProcess (appProc)

    err := appProc.Process.Kill(); if err == nil {

    } else {
        // todo: notify caller about failed operation
    }
    // todo: a. backup app's info
    //       b. kill process with key 'app'
    //       c. rebuild app
    //       d. relaunch app
    return nil
}


func BackupProcess (command *exec.Cmd) ProcessInfo {
    return ProcessInfo { command.Path, command.Args, command.Process.Pid }
}


// vim: noai:ts=4:sw=4
