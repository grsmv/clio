package helpers

import (
	"regexp"
	"strings"
)

/**
 *  Creating a Regexp from custom-formatted pattern
 *  Example:
 *    /users/:id -> "^/users/(?P<id>[\p{L}\d-_]{1,})$"
 */
func PreparePattern(tracery string) *regexp.Regexp {
	tracery = strings.Replace(tracery, ".", "\\.", -1)

	// detecting keywords to convert
	ptrn, _ := regexp.Compile(":[\\p{L}\\d-_]{1,}")

	// converting keywords to regexp form
	ptrn.ReplaceAllStringFunc(tracery, func(match string) string {
		key := strings.Replace(match, ":", "", -1)
		tracery = strings.Replace(tracery, match, "(?P<"+key+">[\\p{L}\\d-_]{1,})", -1)
		return match
	})

	pattern, _ := regexp.Compile("^" + tracery + "$")
	return pattern
}

/**
 *  Converting path to a key-value storage
 */
func ParseSplat(pattern *regexp.Regexp, path string) map[string]string {
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
 *  Splitting path into two parts - absolute page path and query
 *  Example:
 *    "/a/b/c?a=b&c=d" should be splitted into
 *    "/a/b/c" and "a=b&c=d"
 */
func SplitPath(path string) (abs, query string) {
	pattern, _ := regexp.Compile("(.*)\\?(.*)")
	parts := pattern.FindAllStringSubmatch(path, 100)
	if len(parts) > 0 {
		abs = parts[0][1]

		if len(parts[0]) > 1 {
			query = parts[0][2]
		} else {
			query = ""
		}
	} else {
		abs, query = path, ""
	}
	return abs, query
}

/**
 *  Parsing query into map
 *  Example:
 *    query string "a=b&c=d" should become
 *    map[string] string {
 *      "a": "b", "c" : "d"
 *    }
 */
func ParseQuery(queryString string) map[string]string {
	query := make(map[string]string)
	if len(queryString) > 0 {
		pairsArray := strings.Split(queryString, "&")

		for index := range pairsArray {
			if strings.Contains(pairsArray[index], "=") {
				pair := strings.Split(pairsArray[index], "=")
				query[pair[0]] = pair[1]
			} else {
				query[pairsArray[index]] = ""
			}
		}
	}
	return query
}

// vim: noai:ts=4:sw=4
