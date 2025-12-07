package application

import (
	"fmt"
	"os"

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

func (cr *CommandRegistry) Execute(cmd *domains.Command) {
	_, redirectArgs := cmd.Args, cmd.RedirectArg

	if redirectArgs.StdOutPath != "" {
		f, err := utils.OpenRedirectFile(redirectArgs.StdOutPath, redirectArgs.StdOutAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open target file failed with error: %v\n", err)
			return
		}
		defer f.Close()
		cmd.Writer = f
	} else if cmd.Writer == nil {
		cmd.Writer = os.Stdout
	}

	if redirectArgs.StdErrPath != "" {
		f, err := utils.OpenRedirectFile(redirectArgs.StdErrPath, redirectArgs.StdOutAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open target file failed with error: %v\n", err)
		}
		defer f.Close()
		cmd.ErrWriter = f
	} else {
		cmd.ErrWriter = os.Stderr
	}

	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	executor, exists := cr.executors[cmd.Name]

	if !exists {
		// Check for built in external program
		if _, externalExists := utils.FindBinaryInPath(cmd.Name); externalExists {
			builtInExecutor, _ := cr.executors["external-built-in"]
			builtInExecutor.Execute(cmd)
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
		// Ignore external built-in command
		if k == "external-built-in" {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}
