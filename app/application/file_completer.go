package application

import (
	"os"
	"path/filepath"
	"sort"
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

	if token == "" {
		token = "."
	}

	if strings.HasSuffix(token, "/") {
		token += "."
	}

	dir := filepath.Dir(token)
	prefix := filepath.Base(token)

	if prefix == "." {
		prefix = ""
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

	if len(matches) == 0 {
		return nil, 0
	}

	var completions [][]rune
	for _, file := range files {
		if after, ok := strings.CutPrefix(file.Name(), prefix); ok {
			suffix := after

			if file.IsDir() {
				suffix += "/"
			}

			completions = append(completions, []rune(suffix))
		}
	}

	sort.Slice(completions, func(i, j int) bool {
		return string(completions[i]) < string(completions[j])
	})

	if len(completions) == 1 {
		if !strings.HasSuffix(string(completions[0]), "/") {
			completions[0] = append(completions[0], ' ')
		}
	}

	return completions, len(prefix)
}
