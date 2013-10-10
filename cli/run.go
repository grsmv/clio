package cli

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "os/exec"
    "strings"
    "sync"
    // "github.com/cliohq/clio/helpers"
    // "log"
    // "time"
)

const Procfile = "config/procfile.yml"

type Process struct {
    name string
    call string
}

type ProcessList struct {
    processes map[string]string
}

var launchedProcessList map[string]*exec.Cmd


func init () {
    launchedProcessList = map[string]*exec.Cmd {}
}


/**
 *  High-level abstraction for running processes
 *  (called from cli.Route())
 */
func Run() {
    if _, err := os.Stat(Procfile); os.IsNotExist(err) {
        fmt.Println(red, "error:", reset, Procfile, "doesn't exist")
        os.Exit(1)
    }

    // building app before `clio run`
    // log.Print ("Building application binary")
    // helpers.Build()

    // time.Sleep(2 * time.Second)

    go func () {
        list := ProcessList { processes: listProcesses () }
        list.spawnAll ()
    } ()

    // debug
    println (len(launchedProcessList))
}


/**
 * Processing "procfiles" file into map. Each member of
 * resulting map has process symbol-name as key and
 * call string as a value.
 *   @return <Map> list of processes
 */
func listProcesses() map[string]string {
    fileContents, _ := ioutil.ReadFile(Procfile)

    if len(fileContents) == 0 {
        fmt.Println (red, "error", reset, "nothing to run")
        os.Exit(1)
    }

    processesArray := strings.Split(string(fileContents), "\n")
    processesMap := make(map[string]string)

    for i := range processesArray {
        if len(processesArray[i]) > 0 {
            _proc := strings.SplitN(processesArray[i], ":", 2)

            // ignoring commented lines and heredoc lines
            if _proc[0][0] != '#' && len(_proc[1]) > 0 {
                for string(_proc[1][0]) == " " {
                    _proc[1] = strings.Replace(_proc[1], " ", "", 1)
                }
                processesMap[_proc[0]] = _proc[1]
            }
        }
    }

    return processesMap
}


/**
 * Spawning undividual process
 */
func (process *Process) spawn () {

    callParts := strings.Split(process.call, " ")

    command := exec.Command(callParts[0], callParts[1:]...)
    stdOut, _ := command.StdoutPipe()
    stdErr, _ := command.StderrPipe()

    // placing process in global-accessible list
    launchedProcessList[callParts[0]] = command;

    err := command.Start()

    if err == nil {
        fmt.Println(green, process.name, reset, "started")
    } else {
        fmt.Println(red, process.name, reset, "error occured:", err)
        os.Exit(1)
    }

    for _, pipe := range []io.ReadCloser {stdErr, stdOut} {
        go func (pipe io.ReadCloser) {
            reader := bufio.NewReaderSize(pipe, 4*1024)
            line, err := reader.ReadString('\n')

            for err == nil {
                line, err = reader.ReadString('\n')
                fmt.Printf("%s %s %s %s", green, process.name, reset, string(line))
            }
        } (pipe)
    }

    command.Wait()
}


/**
 *  Walking through processe's map and spawning
 *  each member of this map
 */
func (list *ProcessList) spawnAll () {
    var wg sync.WaitGroup

    for name, call := range list.processes {
        wg.Add(1)

        go func (n, c string) {
            process := Process { name: n, call: c }
            process.spawn ()
            wg.Done()
        } (name, call)
    }

    wg.Wait()
}

// vim: noai:ts=4:sw=4
