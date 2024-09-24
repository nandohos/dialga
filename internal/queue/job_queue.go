package queue

import (
	"sync"
	"time"
)

type JobQueue struct {
	Jobs []Job
	Mu   sync.Mutex
}

func NewJobQueue() *JobQueue {
	return &JobQueue{
		Jobs: make([]Job, 0),
	}
}

func (q *JobQueue) EnqueueJob(job Job) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	q.Jobs = append(q.Jobs, job)
}

func (q *JobQueue) DequeueJob() (Job, bool) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if len(q.Jobs) == 0 {
		return Job{}, false
	}
	job := q.Jobs[0]
	q.Jobs = q.Jobs[1:]
	return job, true
}

func (q *JobQueue) GetDueJobs(now time.Time) []Job {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	var dueJobs []Job
	remainingJobs := make([]Job, 0, len(q.Jobs))
	for _, job := range q.Jobs {
		if job.Schedule.Before(now) || job.Schedule.Equal(now) {
			dueJobs = append(dueJobs, job)
		} else {
			remainingJobs = append(remainingJobs, job)
		}
	}
	q.Jobs = remainingJobs
	return dueJobs
}

func (q *JobQueue) Len() int {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	return len(q.Jobs)
}

