package application

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type bellCompleter struct {
	inner readline.AutoCompleter

	tabCount int
}

func (b *bellCompleter) Do(line []rune, pos int) ([][]rune, int) {
	items, length := b.inner.Do(line, pos)

	if len(items) > 1 {
		b.tabCount += 1
		switch b.tabCount {
		case 1:
			lcp := utils.LongestCommonPrefix(items)
			if len(lcp) > 0 {
				b.tabCount = 0
				return [][]rune{lcp}, length
			}

			fmt.Print("\a")

			return nil, 0
		case 2:
			fmt.Println()
			for i, m := range items {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Printf("%s%s", string(line), string(m))
			}

			fmt.Println()

			fmt.Print("\r\x1b[2K")
			fmt.Print("$ ")
			fmt.Print(string(line))

			return nil, 0
		}

	} else if len(items) == 0 {
		fmt.Print("\a")
		return nil, 0
	}

	b.tabCount = 0

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
	var allCmds = append(externalBinaries, builtins...)
	allCmds = utils.DedupeStrings(allCmds)
	baseCompleter := readline.NewPrefixCompleter(readline.PcItemDynamic(func(string) []string {
		return allCmds
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
