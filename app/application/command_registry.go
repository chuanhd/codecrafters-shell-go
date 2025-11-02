package application

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/domains"
)

type CommandRegistry struct {
	executors map[string]domains.CommandExecutor
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		executors: make(map[string]domains.CommandExecutor),
	}
}

func (cr *CommandRegistry) Register(executor domains.CommandExecutor) {
	name := executor.GetName()
	// Throw error if a command has already registered
	if _, exists := cr.executors[name]; exists {
		panic("Command already registered: " + name)
	}

	cr.executors[name] = executor
}

func (cr *CommandRegistry) Execute(cmd domains.Command) {
	executor, exists := cr.executors[cmd.Name]
	if !exists {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.Name)
		// fmt.Fprintf(os.Stdout, "invalid_raspberry_command: command not found")
	} else {
		executor.Execute(cmd)
	}
}
