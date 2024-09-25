package queue

import (
	"sync"
	"testing"
)

func TestNewItemQueue(t *testing.T) {
	name := "TestQueue"
	description := "Test Queue Description"
	queue := NewItemQueue(name, description)

	if queue.Name != name {
		t.Errorf("Expected queue name to be %s, got %s", name, queue.Name)
	}

	if queue.Description != description {
		t.Errorf("Expected queue description to be %s, got %s", description, queue.Description)
	}

	if len(queue.Items) != 0 {
		t.Errorf("Expected queue to be empty, got %d items", len(queue.Items))
	}
}

func TestIncrementCounters(t *testing.T) {
	queue := NewItemQueue("TestQueue", "Test Queue Description")

	queue.IncrementSuccessful()
	if queue.SuccessCount != 1 {
		t.Errorf("Expected SuccessCount to be 1, got %d", queue.SuccessCount)
	}

	queue.IncrementAppException()
	if queue.AppExceptionCount != 1 {
		t.Errorf("Expected AppExceptionCount to be 1, got %d", queue.AppExceptionCount)
	}

	queue.IncrementBizException()
	if queue.BizExceptionCount != 1 {
		t.Errorf("Expected BizExceptionCount to be 1, got %d", queue.BizExceptionCount)
	}
}

func TestEnqueueItem(t *testing.T) {
	queue := NewItemQueue("TestQueue", "Test Queue Description")
	item := Item{} // Assuming Item is defined elsewhere

	queue.EnqueueItem(item)

	if len(queue.Items) != 1 {
		t.Errorf("Expected queue to have 1 item, got %d items", len(queue.Items))
	}
}

func TestConcurrentEnqueueItem(t *testing.T) {
	queue := NewItemQueue("TestQueue", "Test Queue Description")
	itemCount := 1000
	var wg sync.WaitGroup

	for i := 0; i < itemCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			item := Item{} // Assuming Item is defined elsewhere
			queue.EnqueueItem(item)
		}()
	}

	wg.Wait()

	if len(queue.Items) != itemCount {
		t.Errorf("Expected queue to have %d items, got %d items", itemCount, len(queue.Items))
	}
}

func TestConcurrentIncrementCounters(t *testing.T) {
	queue := NewItemQueue("TestQueue", "Test Queue Description")
	incrementCount := 1000
	var wg sync.WaitGroup

	for i := 0; i < incrementCount; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			queue.IncrementSuccessful()
		}()
		go func() {
			defer wg.Done()
			queue.IncrementAppException()
		}()
		go func() {
			defer wg.Done()
			queue.IncrementBizException()
		}()
	}

	wg.Wait()

	if queue.SuccessCount != int64(incrementCount) {
		t.Errorf("Expected SuccessCount to be %d, got %d", incrementCount, queue.SuccessCount)
	}
	if queue.AppExceptionCount != int64(incrementCount) {
		t.Errorf("Expected AppExceptionCount to be %d, got %d", incrementCount, queue.AppExceptionCount)
	}
	if queue.BizExceptionCount != int64(incrementCount) {
		t.Errorf("Expected BizExceptionCount to be %d, got %d", incrementCount, queue.BizExceptionCount)
	}
}
