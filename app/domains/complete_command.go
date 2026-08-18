package domains

import (
	"fmt"
)

type CompleteCommand struct{}

func NewCompleteCommand() *CompleteCommand {
	return &CompleteCommand{}
}

func (c *CompleteCommand) GetName() string {
	return "complete"
}

func (c *CompleteCommand) Execute(cmd *Command) (*ExecuteResult, error) {
	if len(cmd.Args) > 1 {
		switch cmd.Args[0] {
		case "-p":
			specCmd := cmd.Args[1]
			errStr := fmt.Sprintf("complete: %s: no completion specification", specCmd)
			fmt.Fprintln(cmd.Writer, errStr)
			return NewRandomExecuteResult(), fmt.Errorf(errStr)
		}
	}

	return NewRandomExecuteResult(), nil
}
