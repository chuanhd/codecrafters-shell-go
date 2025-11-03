package domains

import (
	"fmt"
	"os"
	"slices"
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

func (c *TypeCommand) Execute(cmd Command) {
	if slices.Contains(c.availableCommands, cmd.Args[0]) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd.Args[0])
	} else {
		fmt.Fprintf(os.Stdout, "%s: not found\n", cmd.Args[0])
	}
}
