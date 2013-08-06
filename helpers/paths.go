package helpers

import (
    "github.com/pallada/clio/vendor/osext"
    "os"
    "strings"
)

/**
 *  Getting application path assuming that application binary placed
 *  in `bin` subdirectory
 */
func ApplicationRoot () string {
    binPath, _ := osext.Executable()
    dirs := strings.Split(binPath, string(os.PathSeparator))
    return strings.Join(dirs[0:(len(dirs) - 2)], string(os.PathSeparator))
}

// vim: noai:ts=4:sw=4
