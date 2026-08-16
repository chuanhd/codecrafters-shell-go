package domains

type CompleteCommand struct{}

func NewCompleteCommand() *CompleteCommand {
	return &CompleteCommand{}
}

func (c *CompleteCommand) GetName() string {
	return "complete"
}

func (c *CompleteCommand) Execute(cmd *Command) (*ExecuteResult, error) {
	return nil, nil
}
