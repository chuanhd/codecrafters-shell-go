package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/application"
	"github.com/codecrafters-io/shell-starter-go/app/domains"
	"github.com/codecrafters-io/shell-starter-go/app/infra"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	handleCommand()
}

func handleCommand() {
	cmdRegistry := application.NewCommandRegistry()
	history := infra.NewInMemoryHistory()

	// Load history on start up
	if histFile := os.Getenv("HISTFILE"); histFile != "" {
		if lines, err := loadHistoryFromFile(histFile); err == nil {
			history.Load(lines)
		}
	}

	cmdHandler := application.NewCommandHandler(cmdRegistry, history)

	// Register the `exit` command
	exitCmd := &domains.ExitCommand{}
	cmdRegistry.Register(exitCmd)

	// Register the `echo` command
	echoCmd := &domains.EchoCommand{}
	cmdRegistry.Register(echoCmd)

	pwdCmd := &domains.PwdCommand{}
	cmdRegistry.Register(pwdCmd)

	cdCmd := &domains.CdCommand{}
	cmdRegistry.Register(cdCmd)

	historyCmd := domains.NewHistoryCommand(history)
	cmdRegistry.Register(historyCmd)

	// Register the `type` command
	// It must be registered last to make sure detect correct supported command
	// Need to append `type` command itself
	supportedCmds := append(cmdRegistry.GetSupportedCmds(), "type")
	typeCmd := domains.NewTypeCommand(supportedCmds)
	cmdRegistry.Register(typeCmd)

	builtInCmd := &domains.ExternalCommand{}
	cmdRegistry.Register(builtInCmd)

	cmdHandler.HandleCommand()
}

func loadHistoryFromFile(path string) ([]string, error) {
	content, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}
