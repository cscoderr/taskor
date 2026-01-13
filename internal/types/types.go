package types

type Worker struct {
	ID int
}

type JobStatus int

const (
	Initiated JobStatus = iota
	Running
	Done
	Cancelled
	Failed
)

func (s JobStatus) String() string {
	switch s {
	case Initiated:
		return "Initiated"
	case Running:
		return "Running"
	case Done:
		return "Done"
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
	Name    string
	Command []string
	Status  JobStatus
}

type Result struct {
	Job    *Job
	Output []byte
	Error  []byte
}
