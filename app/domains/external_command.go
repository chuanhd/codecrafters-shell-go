package domains

import (
	"os/exec"
)

type ExternalCommand struct{}

func (c *ExternalCommand) GetName() string {
	return "external-built-in"
}

func (c *ExternalCommand) Execute(cmd *Command) error {
	externalCmd := exec.Command(cmd.Name, cmd.Args...)
	externalCmd.Stdout = cmd.Writer
	externalCmd.Stderr = cmd.ErrWriter
	externalCmd.Stdin = cmd.Stdin

	return externalCmd.Run()
}
