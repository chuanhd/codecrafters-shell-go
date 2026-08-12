package domains

import (
	"io"
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

type CommandExecutor interface {
	GetName() string
	Execute(cmd *Command) error
}

func (c *Command) IsBackgroundCommand() bool {
	return len(c.Args) > 0 && c.Args[len(c.Args)-1] == "&"
}
