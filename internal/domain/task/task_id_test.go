package task

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTaskID_ShouldGenerateNonEmptyID(t *testing.T) {
	id := NewTaskID()

	if id.value == uuid.Nil {
		t.Fatal("expected a valid UUID, got uuid.Nil")
	}
}

