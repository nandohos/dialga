package queue

import (
	"context"
	"time"
)

// Item represents an item to be processed
type Item struct {
	ID        string
	Priority  int
	Data      interface{}
	CreatedAt time.Time
	StartedAt time.Time
	EndedAt   time.Time
}

type JobStatus string

const (
	StatusPending    JobStatus = "Pending"
	StatusRunning    JobStatus = "Running"
	StatusStopping   JobStatus = "Stopping"
	StatusFaulted    JobStatus = "Faulted"
	StatusSuccessful JobStatus = "Successful"
)

type Process interface {
	Execute(ctx context.Context, job *Job) error
	Name() string
}

// Job represents a scheduled job to process items
type Job struct {
	ID        string
	QueueName string
	Process   Process
	Priority  int
	Schedule  time.Time
	StartedAt time.Time
	EndedAt   time.Time
	State     JobStatus
}
