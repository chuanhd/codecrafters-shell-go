package domains

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type JobsCommand struct {
	jobsList infra.JobsListingsStore
}

func NewJobsCommand(jobsList infra.JobsListingsStore) *JobsCommand {
	return &JobsCommand{
		jobsList: jobsList,
	}
}

func (c *JobsCommand) GetName() string {
	return "jobs"
}

func (c *JobsCommand) Execute(cmd *Command) (*ExecuteResult, error) {

	jobsInBackground := c.jobsList.List()

	if len(jobsInBackground) == 0 {
		return NewRandomExecuteResult(), nil
	}

	fmt.Println(utils.FormatJobsList(
		jobsInBackground,
	))

	c.jobsList.RemoveDoneJobs()

	return NewRandomExecuteResult(), nil
}
