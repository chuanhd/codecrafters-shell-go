package application

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/domains"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
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
		// Check for built in external program
		if path, externalExists := utils.FindBinaryInPath(cmd.Name); externalExists {
			externalCmd := exec.Command(cmd.Name, cmd.Args...)
			output, execErr := externalCmd.Output()
			if execErr != nil {
				fmt.Fprintf(os.Stderr, "Error executing file at %s: %v", path, cmd.Args)
			} else {
				fmt.Fprint(os.Stdout, string(output))
			}
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.Name)
		}

	} else {
		executor.Execute(cmd)
	}
}

func (cr *CommandRegistry) GetSupportedCmds() []string {
	var keys []string
	for k := range cr.executors {
		keys = append(keys, k)
	}
	return keys
}
