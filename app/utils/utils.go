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
