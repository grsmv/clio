package helpers

import (
  "os"
  "strings"
  "strconv"
  "github.com/cliohq/clio/vendor/osext"
  "log"
)

func UpdatePidFile (pid int) {
    pidFile, err := os.Create (pidFilePath ())
    if err != nil {
        log.Fatal (err)
    }
    pidFile.WriteString (strconv.Itoa(pid))
}


func pidFilePath () (path string) {
    bin, _ := osext.Executable ()
    folders := strings.Split (bin, string(os.PathSeparator))
    path = strings.Join (
        folders[0:len(folders) - 1],
        string(os.PathSeparator)) + "/tmp/pids/clio.pid"
    return
}

// vim: noai:ts=4:sw=4
