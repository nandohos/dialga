package queue

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type JobRunner struct {
	jobs   map[string]*Job
	jobsMu sync.RWMutex
	wg     sync.WaitGroup
}

func NewJobRunner() *JobRunner {
	return &JobRunner{
		jobs: make(map[string]*Job),
	}
}

func (jr *JobRunner) AddJob(job *Job) {
	jr.jobsMu.Lock()
	defer jr.jobsMu.Unlock()
	jr.jobs[job.ID] = job
}

func (jr *JobRunner) StartJobs(jobIDs ...string) error {
	for _, jobID := range jobIDs {
		jr.jobsMu.RLock()
		job, exists := jr.jobs[jobID]
		jr.jobsMu.RUnlock()

		if !exists {
			return fmt.Errorf("job with ID %s not found", jobID)
		}

		if job.State == StatusRunning {
			return fmt.Errorf("job %s is already running", jobID)
		}

		job.State = StatusRunning
		job.StartedAt = time.Now()

		ctx, cancel := context.WithCancel(context.Background())

		jr.wg.Add(1)
		go func(j *Job) {
			defer jr.wg.Done()
			defer cancel()

			err := j.Process.Execute(ctx, j)
			j.EndedAt = time.Now()
			if err != nil {
				j.State = StatusFaulted
			} else {
				j.State = StatusSuccessful
			}
		}(job)
	}

	return nil
}

func (jr *JobRunner) StopJob(jobID string) error {
	jr.jobsMu.RLock()
	job, exists := jr.jobs[jobID]
	jr.jobsMu.RUnlock()

	if !exists {
		return fmt.Errorf("job with ID %s not found", jobID)
	}

	if job.State != StatusRunning {
		return fmt.Errorf("job %s is not running", jobID)
	}

	job.State = StatusStopping
	// The actual stopping is handled by canceling the context in the Execute method

	return nil
}

func (jr *JobRunner) WaitForJobs() {
	jr.wg.Wait()
}
