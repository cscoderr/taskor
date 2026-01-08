package engine

import (
	"bytes"
	"context"
	"log"
	"os/exec"
)

type TaskResult struct {
	Output   string
	ExitCode int
	Error    error
}

func ExecuteTask(ctx context.Context, command string) TaskResult {
	log.Println("Start Excecuting Task...")
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	log.Println("Done Excecuting Task...")
	return TaskResult{
		Output: out.String(),
		Error:  err,
	}
}
