package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cscoderr/taskor/internal/engine"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	log.Println("Tasktor Agent Starting...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := engine.ExecuteTask(ctx, "echo 'Hello from Taskor!' && sleep 2 && date")

	if result.Error != nil {
		fmt.Printf("Task Failed: %v\n", result.Error)
	} else {
		fmt.Printf("Task Output:\n%s\n", result.Output)
	}

	//WORKER POOL TEST
	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numOfJobs; a++ {
		<-results
	}
}
