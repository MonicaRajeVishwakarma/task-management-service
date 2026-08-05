package task

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTaskID_ShouldGenerateNonZeroTaskID(t *testing.T) {
	t.Parallel()

	id := NewTaskID()

	if id.value == uuid.Nil {
		t.Fatal("expected a non-zero TaskID")
	}
}

func TestNewTaskID_ShouldGenerateUniqueTaskIDs(t *testing.T) {
	t.Parallel()

	first := NewTaskID()
	second := NewTaskID()

	if first.Equals(second) {
		t.Fatal("expected generated TaskIDs to be unique")
	}
}

func TestTaskID_String_ShouldReturnUUIDString(t *testing.T) {
	t.Parallel()

	id := NewTaskID()

	if id.String() == "" {
		t.Fatal("expected String() to return a UUID")
	}

	if _, err := uuid.Parse(id.String()); err != nil {
		t.Fatalf("expected a valid UUID string, got %q", id.String())
	}
}

func TestTaskID_Equals(t *testing.T) {
	t.Parallel()

	id := NewTaskID()

	if !id.Equals(id) {
		t.Fatal("expected TaskID to equal itself")
	}

	another := NewTaskID()

	if id.Equals(another) {
		t.Fatal("expected different TaskIDs to not be equal")
	}
}