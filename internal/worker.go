package internal

import "math/rand"

type Worker struct {
	ID int
}

func CreateRandWorker() *Worker {
	return &Worker{
		ID: rand.Intn(100),
	}
}
