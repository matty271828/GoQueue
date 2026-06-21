package main

import "fmt"

const workers = 5

func main() {
	ch := make(chan string)

	// launch jobs
	for i := range workers {
		id := fmt.Sprintf("Job %d", i)
		go TestJob(id, ch)
	}

	// receive messages
	for range workers {
		msg := <-ch
		fmt.Println(msg)
	}
}
