package internal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
)

type JobStatus int

const (
	Initiated JobStatus = iota
	Started
	Running
	Done
	Cancelled
	Failed
)

func (s JobStatus) String() string {
	switch s {
	case Initiated:
		return "Initiated"
	case Started:
		return "Started"
	case Done:
		return "Done"
	case Running:
		return "Running"
	case Cancelled:
		return "Cancelled"
	case Failed:
		return "Failed"
	default:
		return "unknown"
	}
}

type Job struct {
	ID      int
	Name    string
	Command []string
	Status  JobStatus
}

type Result struct {
	Job    *Job
	Output []byte
	Error  []byte
}

func CreateJobsFromMap(data map[string]any) []*Job {
	var jobs = make([]*Job, 0, len(data))
	jobId := 1
	for key, value := range data {
		cmd, ok := value.(string)
		if !ok {
			continue
		}
		job := &Job{
			ID:      jobId,
			Name:    key,
			Command: []string{"-c", cmd},
			Status:  Initiated,
		}
		jobs = append(jobs, job)
		jobId++
	}
	return jobs
}

func JobRunner(ctx context.Context, worker *Worker, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result := ExecuteTask(ctx, &job)
		results <- result
	}
}

func ExecuteTask(ctx context.Context, job *Job) Result {
	cmd := exec.CommandContext(ctx, "sh", job.Command...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		job.Status = Failed
		log.Fatalf("could not run command: %v", err)
	}
	job.Status = Done
	return Result{
		Job:    job,
		Output: outBuf.Bytes(),
		Error:  errBuf.Bytes(),
	}
}

func DispatchJobs(jobs []*Job, jobsch chan<- Job) {
	for _, job := range jobs {
		jobsch <- *job
	}
	close(jobsch)
}

func PrintJobResults(results <-chan Result, jobLength int) {
	for result := range results {
		fmt.Println(strings.Repeat("=", 70))
		fmt.Printf("JOB #%d: %s\n", result.Job.ID, result.Job.Name)
		fmt.Printf("COMMAND: %v\n", result.Job.Command)
		fmt.Println(strings.Repeat("-", 70))
		fmt.Printf("Status : %s\n", result.Job.Status)
		if string(result.Output) != "" {
			fmt.Println("Output :")
			fmt.Println(string(result.Output))
		}
		if string(result.Error) != "" {
			fmt.Println("Error :")
			fmt.Println(string(result.Error))
		}
		fmt.Println(strings.Repeat("=", 70))
		fmt.Println()
	}
}
