package internal

type JobStatus int

const (
	Started JobStatus = iota
	Running
	Done
	Cancelled
	Failed
)

func (s JobStatus) String() string {
	switch s {
	case Started:
		return "Started"
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
	Command string
	Status  JobStatus
}

type Result struct {
	JobID  int
	status JobStatus
}

func JobRunner(jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		///run job command
		results <- Result{
			JobID:  job.ID,
			status: Done,
		}
	}
}
