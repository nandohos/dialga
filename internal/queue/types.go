package queue

import "time"

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
	VeryHigh
)

// Item represents an item to be processed
type Item struct {
	ID        string
	Priority  Priority
	Data      interface{}
	CreatedAt time.Time
	StartedAt time.Time
	EndedAt   time.Time
}

// Job represents a scheduled job to process items
type Job struct {
	ID          string
	QueueName   string
	ProcessName string
	Schedule    time.Time
	CreatedAt   time.Time
}
