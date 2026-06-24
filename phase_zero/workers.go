package phase_zero

import (
	"fmt"
)

const workers = 5

// SimpleWorkers is a demonstration of printing
// job ids using basic parallel workers and also
// sequential print statements
func SimpleWorkers() {
	fmt.Println("----- Starting Parallel Execution -----")
	parallelJobs := parallel()
	for _, job := range parallelJobs {
		fmt.Println(job)
	}

	fmt.Println("----- Starting Sequential Execution")
	sequentialJobs := sequential()
	for _, job := range sequentialJobs {
		fmt.Println(job)
	}
}

// launch jobs and receive messages in parallel
func parallel() []string {
	var jobs []string
	ch := make(chan string)

	// launch jobs
	for i := range workers {
		id := fmt.Sprintf("Job %d", i)
		go TestJob(id, ch)
	}

	// receive messages
	for range workers {
		msg := <-ch
		jobs = append(jobs, msg)
	}
	return jobs
}

// launch jobs and receive messages sequentially
func sequential() []string {
	var jobs []string

	for i := range workers {
		id := fmt.Sprintf("Job %d", i)
		jobs = append(jobs, TestSequentialJob(id))
	}

	return jobs
}
