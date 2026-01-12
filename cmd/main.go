package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cscoderr/taskor/internal"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		cancel()
	}()

	jobsJson, numOfWorkers := internal.InputFlagHandler()

	output := internal.ParseJobJson(jobsJson)

	var workers = internal.NewWorkers(*numOfWorkers)

	var jobs = internal.CreateJobsFromMap(output)
	jobsch := make(chan internal.Job, len(jobs))
	resultsch := make(chan internal.Result, len(jobs))

	internal.StartWorkers(ctx, workers, jobsch, resultsch, &wg)
	internal.DispatchJobs(jobs, jobsch)

	go func() {
		wg.Wait()
		close(resultsch)
	}()
	internal.PrintJobResults(resultsch, len(jobs))

	<-ctx.Done()
	fmt.Println("Shut down gracefully!!!")
}
