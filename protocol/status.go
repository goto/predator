package protocol

import (
	"github.com/goto/predator/protocol/job"
	"time"
)

//Status is status of any task
type Status struct {
	ID             string
	JobID          string
	JobType        job.Type
	Status         string
	Message        string
	EventTimestamp time.Time
}
