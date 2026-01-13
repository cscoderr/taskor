package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"runtime"
)

func InputFlagHandler() (*string, *int) {
	jobsJson := flag.String("jobs", "", "List of jobs to run")

	numOfWorkers := flag.Int("workers", runtime.NumCPU(), "Number of concurrent workers")
	flag.Parse()

	if *jobsJson == "" {
		log.Fatalf("Please provide a jobs to run with -job flag")
	}
	return jobsJson, numOfWorkers
}

func ParseJobJson(jobsJson *string) map[string]any {
	jobsByte, err := os.ReadFile(*jobsJson)

	if err != nil {
		log.Fatalf("Unable to read file %s", *jobsJson)
	}

	var output map[string]any
	err = json.Unmarshal(jobsByte, &output)

	if err != nil {
		log.Fatalf("Unable to parse the input, Try again")
	}
	return output
}
