package domains

import (
	"io"
)

type Command struct {
	Name        string
	Args        []string
	RedirectArg RedirectArgument

	Stdin     io.Reader
	Writer    io.Writer
	ErrWriter io.Writer
}

type CommandExecutor interface {
	GetName() string
	Execute(cmd *Command)
}
