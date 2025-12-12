package utils

import (
	"errors"
	"io/fs"
	"os"
	"slices"
	"strings"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ListAllBinariesInPath() []string {
	paths := os.Getenv("PATH")
	var binaries []string
	for path := range strings.SplitSeq(paths, ":") {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		infos := make([]fs.FileInfo, 0, len(entries))
		for _, e := range entries {
			info, _ := e.Info()
			infos = append(infos, info)
		}

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

func OpenFile(path string, needAppend bool) (*os.File, error) {
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

func DedupeStrings(items []string) []string {
	seen := make(map[string]struct{})
	for _, it := range items {
		seen[it] = struct{}{}
	}

	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}

	slices.Sort(out)
	return out
}

func LongestCommonPrefix(items [][]rune) []rune {
	if len(items) == 0 {
		return []rune("")
	}

	if len(items) == 1 {
		return items[0]
	}

	minLength := len(items[0])

	for _, element := range items {
		if len(element) < minLength {
			minLength = len(element)
		}
	}

	for i := range minLength {
		for _, element := range items {
			if element[i] != items[0][i] {
				return element[:i]
			}
		}
	}

	return items[0]
}
