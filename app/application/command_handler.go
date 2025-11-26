package application

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type bellCompleter struct {
	inner readline.AutoCompleter
}

func (b *bellCompleter) Do(line []rune, pos int) ([][]rune, int) {
	items, length := b.inner.Do(line, pos)
	if len(items) == 0 {
		fmt.Print("\a")
	}
	return items, length
}

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
	baseCompleter := readline.NewPrefixCompleter(readline.PcItemDynamic(func(string) []string {
		return append(builtins, externalBinaries...)
	}))

	completer := &bellCompleter{
		inner: baseCompleter,
	}

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
