package utils

import (
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
