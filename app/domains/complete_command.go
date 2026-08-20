package domains

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
)

type CompleteCommand struct {
	completionsStore infra.CompletionsStore
}

func NewCompleteCommand(store infra.CompletionsStore) *CompleteCommand {
	return &CompleteCommand{
		completionsStore: store,
	}
}

func (c *CompleteCommand) GetName() string {
	return "complete"
}

func (c *CompleteCommand) Execute(cmd *Command) (*ExecuteResult, error) {
	if len(cmd.Args) > 1 {
		switch cmd.Args[0] {
		case "-p":
			specCmd := cmd.Args[1]
			script, err := c.completionsStore.Get(specCmd)
			if err != nil {
				fmt.Fprintln(cmd.ErrWriter, err)
				return NewRandomExecuteResult(), err
			}
			fmt.Fprintf(cmd.Writer, "complete -C '%s' %s\n", script, specCmd)
			return NewRandomExecuteResult(), nil
		case "-C":
			if len(cmd.Args) < 3 {
				err := errors.New("complete: -C requires a completer script and a command")
				fmt.Fprintln(cmd.ErrWriter, err)
				return NewRandomExecuteResult(), err
			}
			c.completionsStore.Set(cmd.Args[2], cmd.Args[1])
			return NewRandomExecuteResult(), nil
		default:
			fmt.Fprintln(cmd.ErrWriter, "Unsupported argument")
		}
	}

	return NewRandomExecuteResult(), nil
}
