package helpers

import (
    "strings"
    "regexp"
)

/**
 *  Creating a Regexp from custom-formatted pattern
 *  Example:
 *    /users/:id -> "^/users/(?P<id>[\p{L}\d-_]{1,})$"
 */
func PreparePattern (tracery string) *regexp.Regexp {
    tracery = strings.Replace (tracery, ".", "\\.", -1)

    // detecting keywords to convert
    ptrn, _ := regexp.Compile(":[\\p{L}\\d-_]{1,}")

    // converting keywords to regexp form
    ptrn.ReplaceAllStringFunc(tracery, func(match string) string {
        key := strings.Replace(match, ":", "", -1)
        tracery = strings.Replace(tracery, match, "(?P<" + key + ">[\\p{L}\\d-_]{1,})", -1)
        return match
    })

    pattern, _ := regexp.Compile ("^" + tracery + "$")
    return pattern
}


/**
 *  Converting path to a key-value storage
 */
func ParseSplat (pattern *regexp.Regexp, path string) map[string]string {
    var (
        vocabulary = make(map[string]string)
        matches    = pattern.FindAllStringSubmatch(path, 100)[0][1:]
    )

    for index, key := range pattern.SubexpNames()[1:] {
        vocabulary[key] = matches[index]
    }

    return vocabulary
}


/**
 *  Splitting path into two parts - absolute page path and params
 *  Example:
 *    "/a/b/c?a=b&c=d" should be splitted into
 *    "/a/b/c" and "a=b&c=d"
 */
func SplitPath (path string) (abs, params string) {
    pattern, _ := regexp.Compile("(.*)\\?(.*)")
    parts := pattern.FindAllStringSubmatch(path, 100)
    if len(parts) > 0 {
        abs = parts[0][1]

        if len(parts[0]) > 1 {
            params = parts[0][2]
        } else {
            params = ""
        }
    } else {
        abs, params = path, ""
    }
    return abs, params
}


/**
 *  Parsing params into map
 *  Example:
 *    params string "a=b&c=d" should become
 *    map[string] string {
 *      "a": "b", "c" : "d"
 *    }
 */
func ParseParams (paramsString string) map[string]string {
    params := make(map[string]string)
    if len(paramsString) > 0 {
        pairsArray := strings.Split(paramsString, "&")

        for index := range pairsArray {
            if strings.Contains(pairsArray[index], "=") {
                pair := strings.Split(pairsArray[index], "=")
                params[pair[0]] = pair[1]
            } else {
                params[pairsArray[index]] = ""
            }
        }
    }
    return params
}

// vim: noai:ts=4:sw=4
