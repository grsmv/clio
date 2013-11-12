package helpers

import "testing"

func TestParseParam(t *testing.T) {
    paramsString := "a=b&c=d"
    parsedParams := ParseParams (paramsString)

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
    parsedParams := ParseParams (paramsString)

    if parsedParams["c"] != "" {
        t.Error ("")
    }
}


func TestSplitPath (t *testing.T) {
    abs, params := SplitPath ("/a/b/c?a=b&c=d")
    if abs != "/a/b/c" {
        t.Error ("")
    }

    if params != "a=b&c=d" {
        t.Error ("")
    }
}


func TestSplitPathWithoutParams (t *testing.T) {
    _, params := SplitPath ("/a/b/c?")
    if params != "" {
        t.Error ("")
    }
}


func TestSplitPathWithEmptyString (t *testing.T) {
    abs, params := SplitPath ("")
    if abs != "" {
        t.Error ("")
    }

    if params != "" {
        t.Error ("")
    }
}


func TestPreparePattern (t *testing.T) {
    rxp := PreparePattern ("/a/*/c")
    if rxp.String () != "^/a/([\\p{L}\\d\\-_]{1,})/c$" {
        t.Error ("")
    }
}


func TestPreparePatternDots (t *testing.T) {
    rxp := PreparePattern ("/a.mp3")
    if rxp.String () != "^/a\\.mp3$" {
        t.Error ("")
    }
}
