package internal

import (
	"context"
	"sync"
)

type Worker struct {
	ID int
}

func NewWorkers(n int) []*Worker {
	workers := make([]*Worker, n)
	for i := range n {
		workers[i] = &Worker{ID: i + 1}
	}
	return workers
}

func StartWorkers(ctx context.Context, workers []*Worker, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for _, worker := range workers {
		wg.Add(1)
		go func(w *Worker) {
			JobRunner(ctx, w, jobs, results, wg)
		}(worker)
	}
}
