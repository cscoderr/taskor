package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cscoderr/taskor/internal/config"
	"github.com/cscoderr/taskor/internal/job"
	"github.com/cscoderr/taskor/internal/pool"
	"github.com/cscoderr/taskor/internal/types"
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

	jobsJson, numOfWorkers := config.InputFlagHandler()

	output := config.ParseJobJson(jobsJson)

	var workers = pool.NewWorkers(*numOfWorkers)

	var jobs = job.CreateJobsFromMap(output)
	jobsch := make(chan types.Job, len(jobs))
	resultsch := make(chan types.Result, len(jobs))

	pool.StartWorkers(ctx, workers, jobsch, resultsch, &wg)
	job.DispatchJobs(jobs, jobsch)

	go func() {
		wg.Wait()
		close(resultsch)
	}()
	job.PrintJobResults(resultsch, len(jobs))

	<-ctx.Done()
	fmt.Println("Shut down gracefully!!!")
}
