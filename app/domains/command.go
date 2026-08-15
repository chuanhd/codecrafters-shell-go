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
	PID  int
	Wait func() error
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
		PID:  rand.Int(),
		Wait: nil,
	}
}

func NewExecuteResult(pid int, waitFunc func() error) *ExecuteResult {
	return &ExecuteResult{
		PID:  pid,
		Wait: waitFunc,
	}
}
