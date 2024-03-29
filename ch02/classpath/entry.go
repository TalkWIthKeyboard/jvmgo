package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string)  {
	if strings.Contains(path, pathListSeparator) {
		// TODO
	}
	if strings.HasPrefix(path, "*") {
		// TODO
	}
	if strings.HasPrefix(path, ".jar") ||
		strings.HasPrefix(path, ".JAR") ||
		strings.HasPrefix(path, ".zip") ||
		strings.HasPrefix(path, ".ZIP") {
		// TODO
	}
}
