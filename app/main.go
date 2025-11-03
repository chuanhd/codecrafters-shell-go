package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/application"
	"github.com/codecrafters-io/shell-starter-go/app/domains"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	handleCommand()
}

func handleCommand() {
	cmdRegistry := application.NewCommandRegistry()
	cmdHandler := application.NewCommandHandler(cmdRegistry)

	// Register the `exit` command
	exitCmd := &domains.ExitCommand{}
	cmdRegistry.Register(exitCmd)

	// Register the `echo` command
	echoCmd := &domains.EchoCommand{}
	cmdRegistry.Register(echoCmd)

	// Register the `type` command
	// It must be registered last to make sure detect correct supported command
	// Need to append `type` command itself
	supporedCmds := append(cmdRegistry.GetSupportedCmds(), "type")
	typeCmd := domains.NewTypeCommand(supporedCmds)
	cmdRegistry.Register(typeCmd)

	cmdHandler.HandleCommand()
}
