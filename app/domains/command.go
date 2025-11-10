package domains

import "io"

type Command struct {
	Name   string
	Args   []string
	Writer io.Writer
}

type CommandExecutor interface {
	GetName() string
	Execute(cmd Command)
}
