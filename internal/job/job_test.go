package job

import (
	"context"
	"testing"

	"github.com/cscoderr/taskor/internal/types"
)

func TestExecuteTaskDone(t *testing.T) {
	ctx := context.Background()
	job := &types.Job{
		ID:      1,
		Name:    "echo test",
		Command: []string{"-c", "echo hello"},
	}
	result := ExecuteTask(ctx, job)
	expectedStatus := types.Done

	if job.Status != expectedStatus {
		t.Fatalf("expected %v, got %v", expectedStatus, job.Status)
	}

	if string(result.Output) == "" {
		t.Fatalf("expected output got empty")
	}
}

func TestExecuteTaskFailed(t *testing.T) {
	ctx := context.Background()
	job := &types.Job{
		ID:      2,
		Name:    "always fail",
		Command: []string{"-c", "exit 1"},
	}
	result := ExecuteTask(ctx, job)
	expectedStatus := types.Failed

	if job.Status != expectedStatus {
		t.Fatalf("expected %v, got %v", expectedStatus, job.Status)
	}

	if len(result.Error) == 0 {
		t.Fatalf("expected error output")
	}

}
