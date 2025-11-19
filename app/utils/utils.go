package utils

import (
	"errors"
	"os"
	"strings"
)

func FindBinaryInPath(cmd string) (string, bool) {
	paths := os.Getenv("PATH")
	for path := range strings.SplitSeq(paths, ":") {
		file := path + "/" + cmd
		if fileInfo, err := os.Stat(file); err == nil {
			mode := fileInfo.Mode()
			if mode&0111 != 0 {
				return file, true
			}
		}
	}

	return "", false
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true // File exists
	}
	if errors.Is(err, os.ErrNotExist) {
		return false // File does not exist
	}

	return false
}

func OpenRedirectFile(path string, needAppend bool) (*os.File, error) {
	flags := os.O_CREATE | os.O_RDWR
	if needAppend {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}
	f, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}
