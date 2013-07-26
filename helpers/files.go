package helpers

import (
  "os"
  "strings"
  "strconv"
)

var application_pid string = "tmp/pids/application.pid"

/**
 *  Creating a pidfile and saving PID of Application
 */
func CreatePidFile () {
    file, _ := os.Create(strings.Join([]string{ApplicationRoot(), application_pid}, string(os.PathSeparator)))
    file.WriteString(strconv.Itoa(os.Getpid()))
}

// vim: noai:ts=4:sw=4
