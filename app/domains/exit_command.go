package domains

import (
	"strconv"
)

type ExitCommand struct{}

func (c *ExitCommand) GetName() string {
	return "exit"
}

func (c *ExitCommand) Execute(cmd *Command) (*ExecuteResult, error) {
	exitCode := 0
	if len(cmd.Args) > 0 {
		if code, err := strconv.Atoi(cmd.Args[0]); err == nil {
			exitCode = code
		}
	}

	return NewRandomExecuteResult(), &ExitRequest{
		Code: exitCode,
	}
}
