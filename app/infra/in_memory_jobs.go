package infra

import (
	"slices"
)

type JobStatus string

const (
	JobStatusRunning JobStatus = "Running"
	JobStatusDone    JobStatus = "Done"
)

type JobItem struct {
	JobNumber  int
	ProcessId  int
	CommandStr string
	Status     JobStatus
}

type InMemoryJobs struct {
	jobs []JobItem
}

type JobsListingsStore interface {
	List() []JobItem
	ListJobsDone() []JobItem
	Add(processId int, cmd string) JobItem
	MarkDone(pid int)
	RemoveDoneJobs()
}

func NewInMemoryJobs() *InMemoryJobs {
	return &InMemoryJobs{
		jobs: make([]JobItem, 0),
	}
}

func (h *InMemoryJobs) Add(processId int, cmd string) JobItem {
	newJobNumber := len(h.jobs) + 1
	newJob := JobItem{
		JobNumber:  newJobNumber,
		ProcessId:  processId,
		CommandStr: cmd,
		Status:     JobStatusRunning,
	}
	h.jobs = append(h.jobs, newJob)

	return newJob
}

func (h *InMemoryJobs) MarkDone(pid int) {
	for i := range h.jobs {
		if h.jobs[i].ProcessId == pid {
			h.jobs[i].Status = JobStatusDone
			break
		}
	}
}

func (h *InMemoryJobs) List() []JobItem {
	return h.jobs
}

func (h *InMemoryJobs) ListJobsDone() []JobItem {
	var doneJobs []JobItem
	for _, job := range h.jobs {
		if job.Status == JobStatusDone {
			doneJobs = append(doneJobs, job)
		}
	}

	return doneJobs
}

func (h *InMemoryJobs) RemoveDoneJobs() {
	h.jobs = slices.DeleteFunc(h.jobs, func(job JobItem) bool {
		return job.Status == JobStatusDone
	})
}
