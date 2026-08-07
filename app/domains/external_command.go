package domains

import (
	"fmt"
	"os"
	"os/exec"
)

type ExternalCommand struct{}

func (c *ExternalCommand) GetName() string {
	return "external-built-in"
}

func (c *ExternalCommand) Execute(cmd *Command) error {
	argsLen := len(cmd.Args)
	// If lastArg is &, this job will run in background
	if argsLen > 1 && cmd.Args[argsLen-1] == "&" {
		args := cmd.Args[:argsLen-1]
		externalCmd := exec.Command(cmd.Name, args...)
		err := externalCmd.Start()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "[1] %d\n", externalCmd.Process.Pid)
		return nil
	}

	externalCmd := exec.Command(cmd.Name, cmd.Args...)
	externalCmd.Stdout = cmd.Writer
	externalCmd.Stderr = cmd.ErrWriter
	externalCmd.Stdin = cmd.Stdin

	return externalCmd.Run()
}
