package job

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/cscoderr/taskor/internal/types"
)

func CreateJobsFromMap(data map[string]any) []*types.Job {
	var jobs = make([]*types.Job, 0, len(data))
	jobId := 1
	for key, value := range data {
		cmd, ok := value.(string)
		if !ok {
			continue
		}
		job := &types.Job{
			ID:      jobId,
			Name:    key,
			Command: []string{"-c", cmd},
			Status:  types.Initiated,
		}
		jobs = append(jobs, job)
		jobId++
	}
	return jobs
}

func JobRunner(ctx context.Context, worker *types.Worker, jobs <-chan types.Job, results chan<- types.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		job.Status = types.Running
		result := ExecuteTask(ctx, &job)
		results <- result
	}
}

func ExecuteTask(ctx context.Context, job *types.Job) types.Result {
	cmd := exec.CommandContext(ctx, "sh", job.Command...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		job.Status = types.Failed
		return types.Result{
			Job:   job,
			Error: []byte(err.Error()),
		}
	}
	job.Status = types.Done
	return types.Result{
		Job:    job,
		Output: outBuf.Bytes(),
		Error:  errBuf.Bytes(),
	}
}

func DispatchJobs(jobs []*types.Job, jobsch chan<- types.Job) {
	for _, job := range jobs {
		jobsch <- *job
	}
	close(jobsch)
}

func PrintJobResults(results <-chan types.Result, jobLength int) {
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
