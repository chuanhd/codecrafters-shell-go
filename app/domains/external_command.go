package domains

import (
	"os/exec"
)

type ExternalCommand struct{}

func (c *ExternalCommand) GetName() string {
	return "external-built-in"
}

func (c *ExternalCommand) Execute(cmd *Command) (*ExecuteResult, error) {
	args := cmd.Args
	isBackground := cmd.IsBackgroundCommand()
	if isBackground {
		args = args[:len(args)-1]
	}

	externalCmd := exec.Command(cmd.Name, args...)
	externalCmd.Stdout = cmd.Writer
	externalCmd.Stderr = cmd.ErrWriter
	externalCmd.Stdin = cmd.Stdin

	if err := externalCmd.Start(); err != nil {
		return nil, err
	}

	if isBackground {
		return NewExecuteResult(externalCmd.Process.Pid, externalCmd.Wait), nil
	}

	return NewExecuteResult(externalCmd.Process.Pid, nil), externalCmd.Wait()
}
