package main

import (
	"fmt"

	"github.com/cscoderr/taskor/internal"
)

func main() {
	var workers = []*internal.Worker{
		internal.CreateRandWorker(),
		internal.CreateRandWorker(),
		internal.CreateRandWorker(),
	}

	fmt.Printf("%v\n", workers)

}
