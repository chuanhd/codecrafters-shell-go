package utils

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

func ListAllBinariesInPath() []string {
	paths := os.Getenv("PATH")
	var binaries []string
	for path := range strings.SplitSeq(paths, ":") {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		infos := make([]fs.FileInfo, 0, len(entries))
		for _, fileInfo := range infos {
			mode := fileInfo.Mode()
			if mode&0111 != 0 {
				binaries = append(binaries, fileInfo.Name())
			}
		}
	}

	return binaries
}

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
