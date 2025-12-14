package domains

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c *CdCommand) GetName() string {
	return "cd"
}

func (c *CdCommand) Execute(cmd *Command) error {
	path := cmd.Args[0]
	if path == "~" {
		homeDir, _err := os.UserHomeDir()
		if _err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get home directory %s\n", _err.Error())
			return _err
		}
		path = homeDir
	}

	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s: No such file or directory\n", cmd.Name, path)
		return err
	}

	err = os.Chdir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory to %s\n", cmd.Args[0])
	}

	return nil
}
