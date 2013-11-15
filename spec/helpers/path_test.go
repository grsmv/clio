package helpers

import (
    "os"
    "testing"
    "github.com/cliohq/clio/helpers"
)

func TestFixPath(t *testing.T) {
    if string (os.PathSeparator) == "/" {
        if helpers.FixPath("/a/b/c") != "/a/b/c" {
            t.Error ("")
        }
    } else {
        if helpers.FixPath ("/a/b/c") != "\\a\\b\\c" {
            t.Error ("")
        }
    }
}
