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

// TaskIDFromString creates a TaskID from an existing UUID string.
func TaskIDFromString(id string) (TaskID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return TaskID{}, ErrInvalidTaskID
	}

	return TaskID{
		value: parsedID,
	}, nil
}

// String returns the string representation of TaskID.
func (id TaskID) String() string {
	return id.value.String()
}

// Equals compares two TaskIDs.
func (id TaskID) Equals(other TaskID) bool {
	return id.value == other.value
}
