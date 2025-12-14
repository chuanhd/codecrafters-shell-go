package domains

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (c *PwdCommand) GetName() string {
	return "pwd"
}

func (c *PwdCommand) Execute(cmd *Command) error {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	fmt.Fprintln(os.Stdout, path)

	return nil
}
