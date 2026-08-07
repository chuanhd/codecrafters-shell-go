package domains

type JobsCommand struct{}

func (c *JobsCommand) GetName() string {
	return "jobs"
}

func (c *JobsCommand) Execute(cmd *Command) error {
	return nil
}
