package domains

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
)

type HistoryCommand struct {
	history infra.HistoryStore
}

func NewHistoryCommand(history infra.HistoryStore) *HistoryCommand {
	return &HistoryCommand{
		history: history,
	}
}

func (cmd *HistoryCommand) GetName() string {
	return "history"
}

func (c *HistoryCommand) Execute(cmd *Command) {
	c.history.Add(cmd.RawContent)

	var total = len(c.history.List())
	var limit = total
	if len(cmd.Args) > 0 {
		if limitArg, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = limitArg
		}
	}
	// Add itself at last
	for i := total - limit; i < total; i++ {
		line := c.history.Get(i)
		fmt.Fprintf(cmd.Writer, "%d  %s\n", i, line)
	}
}
