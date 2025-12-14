package domains

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type TypeCommand struct {
	availableCommands []string
}

func NewTypeCommand(availableCmds []string) *TypeCommand {
	return &TypeCommand{
		availableCommands: availableCmds,
	}
}

func (c *TypeCommand) GetName() string {
	return "type"
}

func (c *TypeCommand) Execute(cmd *Command) error {
	if slices.Contains(c.availableCommands, cmd.Args[0]) {
		fmt.Fprintf(cmd.Writer, "%s is a shell builtin\n", cmd.Args[0])
	} else if path, exists := c.findBinInPath(*cmd); exists {
		fmt.Fprintf(cmd.Writer, "%s is %s\n", cmd.Args[0], path)
	} else {
		fmt.Fprintf(cmd.Writer, "%s: not found\n", cmd.Args[0])
	}

	return nil
}

func (c *TypeCommand) findBinInPath(cmd Command) (string, bool) {
	paths := os.Getenv("PATH")
	bin := cmd.Args[0]
	for path := range strings.SplitSeq(paths, ":") {
		file := path + "/" + bin
		if fileInfo, err := os.Stat(file); err == nil {
			mode := fileInfo.Mode()
			if mode&0111 != 0 {
				return file, true
			}
		}
	}

	return "", false
}
