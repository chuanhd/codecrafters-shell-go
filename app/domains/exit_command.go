package domains

import (
	"fmt"
	"os"
	"strconv"
)

type ExitCommand struct{}

func (c *ExitCommand) GetName() string {
	return "exit"
}

func (c *ExitCommand) Execute(cmd Command) {
	exitCode, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		fmt.Println("Error parsing exit code: ", err)
	}
	os.Exit(exitCode)
}
