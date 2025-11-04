package domains

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c *CdCommand) GetName() string {
	return "cd"
}

func (c *CdCommand) Execute(cmd Command) {
	path := cmd.Args[0]
	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s: No such file or directory\n", cmd.Name, path)
		return
	}
	err = os.Chdir(cmd.Args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory to %s\n", cmd.Args[0])
	}
}
