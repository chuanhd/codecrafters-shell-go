package application

import (
	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
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
	var externalBinaries = utils.ListAllBinariesInPath()
	var builtins = ch.registry.GetSupportedCmds()
	completer := readline.NewPrefixCompleter(readline.PcItemDynamic(func(string) []string {
		return append(builtins, externalBinaries...)
	}))

	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "$ ",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	parser := NewCommandParser(rl)

	for {
		cmd, err := parser.ParseCommand()
		if err != nil {
			continue
		}

		ch.registry.Execute(cmd)
	}
}
