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
    for proc, _ := range LaunchedProcesses {
        print (proc)
    }

    // todo: a. backup app's info
    //       b. kill process with key 'app'
    //       c. rebuild app
    //       d. relaunch app
    return nil
}

// vim: noai:ts=4:sw=4
