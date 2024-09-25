package queue

import (
	"testing"
	"time"
)

func TestNewJobQueue(t *testing.T) {
	q := NewJobQueue()
	if q == nil {
		t.Fatal("NewJobQueue returned nil")
	}
	if len(q.Jobs) != 0 {
		t.Errorf("Expected empty queue, got %d jobs", len(q.Jobs))
	}
}

func TestEnqueueJob(t *testing.T) {
	q := NewJobQueue()
	job := Job{ID: "1", Schedule: time.Now()}
	q.EnqueueJob(job)
	if len(q.Jobs) != 1 {
		t.Errorf("Expected 1 job, got %d", len(q.Jobs))
	}
	if q.Jobs[0] != job {
		t.Errorf("Expected job %v, got %v", job, q.Jobs[0])
	}
}

func TestDequeueJob(t *testing.T) {
	q := NewJobQueue()
	job := Job{ID: "1", Schedule: time.Now()}
	q.EnqueueJob(job)

	dequeuedJob, ok := q.DequeueJob()
	if !ok {
		t.Fatal("DequeueJob returned false, expected true")
	}
	if dequeuedJob != job {
		t.Errorf("Expected job %v, got %v", job, dequeuedJob)
	}
	if len(q.Jobs) != 0 {
		t.Errorf("Expected empty queue, got %d jobs", len(q.Jobs))
	}

	_, ok = q.DequeueJob()
	if ok {
		t.Error("DequeueJob returned true for empty queue, expected false")
	}
}

func TestGetDueJobs(t *testing.T) {
	q := NewJobQueue()
	now := time.Now()
	pastJob := Job{ID: "1", Schedule: now.Add(-1 * time.Hour)}
	futureJob := Job{ID: "2", Schedule: now.Add(1 * time.Hour)}
	q.EnqueueJob(pastJob)
	q.EnqueueJob(futureJob)

	dueJobs := q.GetDueJobs(now)
	if len(dueJobs) != 1 {
		t.Errorf("Expected 1 due job, got %d", len(dueJobs))
	}
	if dueJobs[0] != pastJob {
		t.Errorf("Expected due job %v, got %v", pastJob, dueJobs[0])
	}
	if len(q.Jobs) != 1 {
		t.Errorf("Expected 1 remaining job, got %d", len(q.Jobs))
	}
	if q.Jobs[0] != futureJob {
		t.Errorf("Expected remaining job %v, got %v", futureJob, q.Jobs[0])
	}
}

func TestLen(t *testing.T) {
	q := NewJobQueue()
	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}

	q.EnqueueJob(Job{ID: "1", Schedule: time.Now()})
	if q.Len() != 1 {
		t.Errorf("Expected length 1, got %d", q.Len())
	}

	q.EnqueueJob(Job{ID: "2", Schedule: time.Now()})
	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}

	q.DequeueJob()
	if q.Len() != 1 {
		t.Errorf("Expected length 1 after dequeue, got %d", q.Len())
	}
}
