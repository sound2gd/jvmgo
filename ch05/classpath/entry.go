package classpath

import (
	"os"
	"strings"
)

const PATH_SEPERATOR = string(os.PathListSeparator)

var JarEntry_Suffix_List = []string{
	".jar",
	".zip",
}

/**
 * golang中的Interface类似于clj的protol
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, PATH_SEPERATOR) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	for _, suffix := range JarEntry_Suffix_List {
		if strings.HasSuffix(strings.ToLower(path), suffix) {
			return newZipEntry(path)
		}

	}
	// if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	// 	strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
	// }
	return newDirEntry(path)
}
