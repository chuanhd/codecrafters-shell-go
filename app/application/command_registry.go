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
	args, outPath, errPath, err := cr.handleStdOut(cmd.Args)
	outWriter := os.Stdout
	errWriter := os.Stderr

	if outPath != "" {
		f, err := os.OpenFile(outPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open target file failed with error: %v\n", err)
			return
		}
		defer f.Close()
		outWriter = f
	}

	if errPath != "" {
		f, err := os.OpenFile(errPath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open target file failed with error: %v\n", err)
		}
		defer f.Close()
		errWriter = f
	}

	cmd.Args = args

	if !exists {
		// Check for built in external program
		if _, externalExists := utils.FindBinaryInPath(cmd.Name); externalExists {
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command failed with error: %v\n", err)
				return
			}
			externalCmd := exec.Command(cmd.Name, args...)
			externalCmd.Stdout = outWriter
			externalCmd.Stderr = errWriter

			externalCmd.Run()
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.Name)
		}

	} else {
		cmd.Writer = outWriter
		cmd.ErrWriter = errWriter
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

func (cr *CommandRegistry) handleStdOut(rawArgs []string) (args []string, outPath string, errPath string, err error) {
	for i := 0; i < len(rawArgs); i++ {
		switch rawArgs[i] {
		case ">", "1>":
			if i+1 > len(rawArgs) {
				return nil, "", "", fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			outPath = rawArgs[i+1]
			i++ // skip filename
		case "2>":
			if i+1 > len(rawArgs) {
				return nil, "", "", fmt.Errorf("syntax error: missing filename after %q", rawArgs[i])
			}
			errPath = rawArgs[i+1]
			i++ // skip filename
		default:
			args = append(args, rawArgs[i])
		}
	}

	return args, outPath, errPath, nil
}
