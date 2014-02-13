package helpers

import (
    "testing"
    "github.com/cliohq/clio/helpers"
)

func TestParseParam(t *testing.T) {
    paramsString := "a=b&c=d"
    parsedParams := helpers.ParseQuery(paramsString)

    if len (parsedParams) != 2 {
        t.Error ("")
    }

    if parsedParams["a"] != "b" {
        t.Error ("")
    }

    if parsedParams["c"] != "d" {
        t.Error ("")
    }
}


func TestParseParamKeyWithoutValue(t *testing.T) {
    paramsString := "a=b&c"
    parsedParams := helpers.ParseQuery(paramsString)

    if parsedParams["c"] != "" {
        t.Error ("")
    }
}


func TestSplitPath(t *testing.T) {
    abs, params := helpers.SplitPath ("/a/b/c?a=b&c=d")
    if abs != "/a/b/c" {
        t.Error ("")
    }

    if params != "a=b&c=d" {
        t.Error ("")
    }
}


func TestSplitPathWithoutParams(t *testing.T) {
    _, params := helpers.SplitPath ("/a/b/c?")
    if params != "" {
        t.Error ("")
    }
}


func TestSplitPathWithEmptyString(t *testing.T) {
    abs, params := helpers.SplitPath("")
    if abs != "" {
        t.Error("")
    }

    if params != "" {
        t.Error("")
    }
}


func TestPreparePattern(t *testing.T) {
    rxp := helpers.PreparePattern("/a/:id/c")
    if rxp.String() != "^/a/(?P<id>[\\p{L}\\d-_]{1,})/c$" {
        t.Error ("")
    }
}


func TestPreparePatternDots(t *testing.T) {
    rxp := helpers.PreparePattern("/a.mp3")
    if rxp.String() != "^/a\\.mp3$" {
        t.Error("")
    }
}
