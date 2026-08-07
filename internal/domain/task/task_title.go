package task

import (
	"errors"
	"strings"
)

// ErrEmptyTaskTitle occurs when a task title is empty.
var ErrEmptyTaskTitle = errors.New("task title can't be empty")

// TaskTitle represents a validated task title.
type TaskTitle struct {
	value string
}

// NewTaskTitle creates a new task title
func NewTaskTitle(title string) (TaskTitle, error) {
	trimmedTitle := strings.TrimSpace(title)
	if trimmedTitle == "" {
		return TaskTitle{}, ErrEmptyTaskTitle
	}

	return TaskTitle{
		value: trimmedTitle,
	}, nil
}

// String returns the value of TaskTitle in string type
func (t TaskTitle) String() string {
	return t.value
}

// Equals return bool value after comparing 2 Tasktitle
func (t TaskTitle) Equals(other TaskTitle) bool {
	return t.value == other.value
}
