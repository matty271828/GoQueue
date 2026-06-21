package main

import (
	"fmt"
	"time"
)

type Job struct {
	id        string
	createdAt time.Time
}

func NewJob(id string, createdAt time.Time) Job {
	return Job{
		id:        id,
		createdAt: createdAt,
	}
}

// TestJob is an example job used to test the job queuer.
func TestJob(id string, ch chan string) {
	job := NewJob(id, time.Now())
	ch <- fmt.Sprintf("%s created at: %s\n", id, job.createdAt)
}
