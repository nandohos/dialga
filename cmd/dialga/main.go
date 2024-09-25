// dialga/cmd/dialga/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nandohos/dialga/internal/queue"
)

type ImAliveProcess struct{}

func (p ImAliveProcess) Name() string {
	return "ImAlive"
}

func (p ImAliveProcess) Execute(ctx context.Context, job *queue.Job) error {
	log.Printf("Starting job: %s\n", job.ID)
	start := time.Now()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			log.Printf("Job %s: I'm Alive\n", job.ID)
			if time.Since(start) >= 60*time.Second {
				return nil
			}
		}
	}
}

type ImAliveMoreOftenProcess struct{}

func (p ImAliveMoreOftenProcess) Name() string {
	return "ImAliveMoreOften"
}

func (p ImAliveMoreOftenProcess) Execute(ctx context.Context, job *queue.Job) error {
	log.Printf("Starting job: %s\n", job.ID)
	start := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			log.Printf("Job %s: I'm Alive and more often\n", job.ID)
			if time.Since(start) >= 60*time.Second {
				return nil
			}
		}
	}
}

func main() {
	fmt.Println("Dialga - Job Scheduler")

	jobRunner := queue.NewJobRunner()

	// Create processes
	imAliveProcess := ImAliveProcess{}
	imAliveMoreOftenProcess := ImAliveMoreOftenProcess{}

	// Create jobs
	job1 := &queue.Job{
		ID:        "job1",
		QueueName: "DefaultQueue",
		Process:   imAliveProcess,
		Priority:  1,
		Schedule:  time.Now(),
		State:     queue.StatusPending,
	}

	job2 := &queue.Job{
		ID:        "job2",
		QueueName: "DefaultQueue",
		Process:   imAliveMoreOftenProcess,
		Priority:  1,
		Schedule:  time.Now(),
		State:     queue.StatusPending,
	}

	jobRunner.AddJob(job1)
	jobRunner.AddJob(job2)

	// Start both jobs simultaneously
	if err := jobRunner.StartJobs("job1", "job2"); err != nil {
		log.Printf("Failed to start jobs: %v", err)
		return
	}

	// Wait for jobs to complete
	jobRunner.WaitForJobs()

	log.Println("All jobs completed")
}
