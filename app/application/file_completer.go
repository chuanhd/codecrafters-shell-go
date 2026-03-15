package application

import (
	"os"
	"path/filepath"
	"strings"
)

func NewFileCompleter() *FileCompleter {
	return &FileCompleter{}
}

type FileCompleter struct {
}

func (f *FileCompleter) Do(line []rune, pos int) ([][]rune, int) {
	input := string(line[:pos])

	parts := strings.Split(input, " ")
	token := parts[len(parts)-1]

	var dir, prefix string
	if before, ok := strings.CutSuffix(token, "/"); ok {
		dir = before
		if dir == "" {
			dir = "."
		}
		prefix = ""
	} else {
		dir = filepath.Dir(token)
		prefix = filepath.Base(token)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0
	}

	var matches []os.DirEntry

	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			matches = append(matches, file)
		}
	}

	var completions [][]rune
	for _, file := range matches {
		name := file.Name()
		suffix := strings.TrimPrefix(name, prefix)
		if len(matches) == 1 {
			if file.IsDir() {
				suffix += "/"
			} else {
				suffix += " "
			}
		}
		completions = append(completions, []rune(suffix))
	}

	return completions, len(prefix)
}
