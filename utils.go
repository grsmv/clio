package hades

import (
    "strings"
    "regexp"
)

/**
 *  Creating a Regexp from custom-formatted pattern
 *  Example:
 *      "/books/ * /update" -> "^/books/([\p{L}\d\-_]{1,})/update$"
 */
func prepearePattern (rawPattern string) *regexp.Regexp {
    replaceRools := strings.NewReplacer(
        "*", "([\\p{L}\\d\\-_]{1,})",
        ".", "\\.")

    // adding start and end line symbols
    pattern := "^" + replaceRools.Replace(rawPattern) + "$"
    regexpPattern, _ := regexp.Compile(pattern)
    return regexpPattern
}


/**
 *  Splitting path into two parts - absolute page path and params
 *  Example:
 *    "/a/b/c?a=b&c=d" should be splitted into
 *    "/a/b/c" and "a=b&c=d"
 */
func splitPath (path string) (abs, params string) {
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
func parseParams (paramsString string) map[string]string {
    params := make(map[string]string)
    pairsArray := strings.Split(paramsString, "&")
    for index := range pairsArray {
        pair := strings.Split(pairsArray[index], "=")
        params[pair[0]] = pair[1]
    }
    return params
}
