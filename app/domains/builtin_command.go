package domains

import (
	"os/exec"
)

type BuiltInCommand struct{}

func (c *BuiltInCommand) GetName() string {
	return "external-built-in"
}

func (c *BuiltInCommand) Execute(cmd *Command) {
	externalCmd := exec.Command(cmd.Name, cmd.Args...)
	externalCmd.Stdout = cmd.Writer
	externalCmd.Stderr = cmd.ErrWriter

	externalCmd.Run()
}
