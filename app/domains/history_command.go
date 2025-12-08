package domains

type HistoryCommand struct{}

func NewHistoryCommand() *HistoryCommand {
	return &HistoryCommand{}
}

func (cmd *HistoryCommand) GetName() string {
	return "history"
}

func (c *HistoryCommand) Execute(cmd *Command) {
	// TODO: Implement history command
}
