package task

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidTaskID = errors.New("invalid task id")

// TaskID represents the unique entity of a task.
type TaskID struct {
	value uuid.UUID
}


// NewTaskID creates a new TaskID
func NewTaskID() TaskID {
	return TaskID{
		value: uuid.New(),
	}
}
