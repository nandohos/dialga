package queue

import (
	"sync"
	"sync/atomic"
)

type ItemQueue struct {
	Name              string
	Description       string
	Items             []Item
	Mutex             sync.Mutex
	SuccessCount      int64
	AppExceptionCount int64
	BizExceptionCount int64
}

// NewItemQueue creates and returns a new ItemQueue with the given name and description
func NewItemQueue(name string, description string) *ItemQueue {
	return &ItemQueue{
		Name:        name,
		Description: description,
		Items:       make([]Item, 0),
	}
}

// IncrementSuccessful atomically increments the SuccessCount
func (q *ItemQueue) IncrementSuccessful() {
	atomic.AddInt64(&q.SuccessCount, 1)
}

// IncrementAppException atomically increments the AppExceptionCount
func (q *ItemQueue) IncrementAppException() {
	atomic.AddInt64(&q.AppExceptionCount, 1)
}

// IncrementBizException atomically increments the BizExceptionCount
func (q *ItemQueue) IncrementBizException() {
	atomic.AddInt64(&q.BizExceptionCount, 1)
}

// EnqueueItem adds an item to the queue in a thread-safe manner
func (q *ItemQueue) EnqueueItem(item Item) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Items = append(q.Items, item)
}
