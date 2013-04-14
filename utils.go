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
