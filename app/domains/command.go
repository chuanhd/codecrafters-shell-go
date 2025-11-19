package domains

import (
	"io"
)

type Command struct {
	Name        string
	Args        []string
	Writer      io.Writer
	ErrWriter   io.Writer
	RedirectArg RedirectArgument
}

type CommandExecutor interface {
	GetName() string
	Execute(cmd *Command)
}
