package queue

import "time"

// Item represents an item to be processed
type Item struct {
	ID        string
	Priority  int
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
	Priority    int
	Schedule    time.Time
	CreatedAt   time.Time
}
