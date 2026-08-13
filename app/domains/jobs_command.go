package domains

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
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

	fmt.Println(formatJobsList(
		jobsInBackground,
	))

	return NewRandomExecuteResult(), nil
}

func formatJobsList(jobs []infra.JobItem) string {
	result := make([]string, 0)

	for index, job := range jobs {
		item := ""
		recentMarker := " "
		if index == len(jobs)-1 {
			recentMarker = "+"
		} else if index == len(jobs)-2 {
			recentMarker = "-"
		}
		item = fmt.Sprintf("[%d]%s  %-24s %s", job.JobNumber, recentMarker, job.Status, job.CommandStr)
		result = append(result, item)
	}

	return strings.Join(result, "\n")
}
