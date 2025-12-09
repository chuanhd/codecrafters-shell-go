package domains

import (
	"fmt"

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
	// Add itself at last
	c.history.Add(cmd.RawContent)
	for i, line := range c.history.List() {
		fmt.Fprintf(cmd.Writer, "%d  %s\n", i+1, line)
	}
}
