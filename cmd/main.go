package main

import (
	"flag"
	"fmt"

	"github.com/cscoderr/taskor/internal"
)

func main() {
	var workers = []*internal.Worker{
		internal.CreateRandWorker(),
		internal.CreateRandWorker(),
		internal.CreateRandWorker(),
	}
	jobs := make(chan internal.Job, 10)
	results := make(chan internal.Result, 10)

	for index, worker := range workers {
		go func() {
			fmt.Printf("Worker index: %d is with ID: %d", index, worker.ID)
		}()
	}

	

	fmt.Printf("%v\n", *workers[2])

	name := flag.String("name", "", "Just a random name")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please provide a name user -name flag")
	}

	fmt.Println(*name)

}
