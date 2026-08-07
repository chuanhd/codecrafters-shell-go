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
	lastArg := cmd.Args[len(cmd.Args)-1]
	// If lastArg is &, this job will run in background
	if lastArg == "&" {
		args := cmd.Args[:len(cmd.Args)]
		fmt.Println(args)
		externalCmd := exec.Command(cmd.Name, args...)
		err := externalCmd.Start()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "[1] %d\n", externalCmd.Process.Pid)
		return nil
	} else {
		externalCmd := exec.Command(cmd.Name, cmd.Args...)
		externalCmd.Stdout = cmd.Writer
		externalCmd.Stderr = cmd.ErrWriter
		externalCmd.Stdin = cmd.Stdin

		return externalCmd.Run()
	}
}
