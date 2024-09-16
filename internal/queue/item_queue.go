package queue

import "sync"

type ItemQueue struct {
	Name              string
	Description       string
	items             []Item
	mu                sync.Mutex
	successCount      int64
	appExceptionCount int64
	bizExceptionCount int64
}
