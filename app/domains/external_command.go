package domains

import (
	"fmt"
	"os/exec"
)

type ExternalCommand struct{}

func (c *ExternalCommand) GetName() string {
	return "external-built-in"
}

func (c *ExternalCommand) Execute(cmd *Command) error {
	args := cmd.Args
	isBackground := len(args) > 0 && args[len(args)-1] == "&"

	if isBackground {
		args = args[:len(args)-1]
	}

	externalCmd := exec.Command(cmd.Name, args...)
	externalCmd.Stdout = cmd.Writer
	externalCmd.Stderr = cmd.ErrWriter
	externalCmd.Stdin = cmd.Stdin

	if err := externalCmd.Start(); err != nil {
		return err
	}

	if isBackground {
		fmt.Fprintf(cmd.Writer, "[1] %d\n", externalCmd.Process.Pid)
		return nil
	}

	return externalCmd.Wait()
}
