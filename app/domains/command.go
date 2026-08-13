package domains

import (
	"io"
	"math/rand"
)

type Command struct {
	Name        string
	Args        []string
	RedirectArg RedirectArgument
	RawContent  string

	Stdin     io.Reader
	Writer    io.Writer
	ErrWriter io.Writer
}

type ExecuteResult struct {
	PID int
}

type CommandExecutor interface {
	GetName() string
	Execute(cmd *Command) (*ExecuteResult, error)
}

func (c *Command) IsBackgroundCommand() bool {
	return len(c.Args) > 0 && c.Args[len(c.Args)-1] == "&"
}

/*
 * Random process id for internal command
 */
func NewRandomExecuteResult() *ExecuteResult {
	return &ExecuteResult{
		PID: rand.Int(),
	}
}

func NewExecuteResult(pid int) *ExecuteResult {
	return &ExecuteResult{
		PID: pid,
	}
}
