package application

import (
	"bufio"
	"fmt"
	"os"
)

type CommandHandler struct {
	registry *CommandRegistry
}

func NewCommandHandler(registry *CommandRegistry) *CommandHandler {
	return &CommandHandler{
		registry: registry,
	}
}

func (ch *CommandHandler) HandleCommand() {
	reader := bufio.NewReader(os.Stdin)
	parser := NewCommandParser(reader)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := parser.ParseCommand()
		if err != nil {
			continue
		}

		ch.registry.Execute(*cmd)

		// reader.Reset(os.Stdin)
	}
}
