package goldflags

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// PathExist checks whether a given path exists or not.
func PathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ExpandPath expands a path name. "~/" is expanded to the current user's
// home directory and "./" prepends the base path (defaults to ".", i.e.
// resulting in a no-op).
// After expansion, the result is cleaned (see filepath.Clean() for details).
func ExpandPath(path string, base ...string) (result string, err error) {
	result = path

	switch path[:2] {
	case "~/":
		var u *user.User
		if u, err = user.Current(); err != nil {
			return
		}
		result = filepath.Join(u.HomeDir, path[2:])
	case "./":
		if len(base) > 0 {
			result = filepath.Join(base[0], path[2:])
		}
	}

	return filepath.Clean(result), nil
}

// ReadDir returns a (non-recursive) list of file and directory names for
// the given directory path.
func ReadDir(dir string) (names []string) {
	if files, err := ioutil.ReadDir(dir); err == nil {
		for _, file := range files {
			names = append(names, file.Name())
		}
	} else {
		log.Print(err)
	}
	return
}
