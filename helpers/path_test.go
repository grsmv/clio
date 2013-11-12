package helpers

import (
    "os"
    "testing"
)

func TestFixPath(t *testing.T) {
    if string (os.PathSeparator) == "/" {
        if FixPath("/a/b/c") != "/a/b/c" {
            t.Error ("")
        }
    } else {
        if FixPath ("/a/b/c") != "\\a\\b\\c" {
            t.Error ("")
        }
    }
}
