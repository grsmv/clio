package helpers

import (
  "os"
  "strings"
  "strconv"
  "log"
)


/**
 *  Creating a pidfile and saving PID of Application
 */
func CreatePidFile (pidFile string) {
    log.Println (ApplicationRoot ())

    file, _ := os.Create(
        strings.Join(
            []string{ ApplicationRoot (), pidFile },
            string(os.PathSeparator)))
    file.WriteString(strconv.Itoa(os.Getpid()))
}

// vim: noai:ts=4:sw=4
