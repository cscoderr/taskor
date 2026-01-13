package pool

import (
	"context"
	"sync"

	"github.com/cscoderr/taskor/internal/job"
	"github.com/cscoderr/taskor/internal/types"
)

func NewWorkers(n int) []*types.Worker {
	workers := make([]*types.Worker, n)
	for i := range n {
		workers[i] = &types.Worker{ID: i + 1}
	}
	return workers
}

func StartWorkers(ctx context.Context, workers []*types.Worker, jobs <-chan types.Job, results chan<- types.Result, wg *sync.WaitGroup) {
	for _, worker := range workers {
		wg.Add(1)
		go func(w *types.Worker) {
			job.JobRunner(ctx, w, jobs, results, wg)
		}(worker)
	}
}
