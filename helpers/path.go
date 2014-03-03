package helpers

import (
	"os"
	"strings"
)

func FixPath(path string) string {
	return strings.Replace(path, "/", string(os.PathSeparator), -1)
}
