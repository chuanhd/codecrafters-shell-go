package domains

import (
	"os"
	"strconv"
)

type ExitCommand struct{}

func (c *ExitCommand) GetName() string {
	return "exit"
}

func (c *ExitCommand) Execute(cmd *Command) {
	if len(cmd.Args) > 0 {
		exitCode, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			exitCode = 0
		}
		os.Exit(exitCode)
	}
	os.Exit(0)
}
