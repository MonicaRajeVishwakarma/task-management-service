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

	if first.value == second.value {
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

	if (id.value != id.value) {
		t.Fatal("expected TaskID to equal itself")
	}

	another := NewTaskID()

	if id.Equals(another) {
		t.Fatal("expected different TaskIDs to not be equal")
	}
}
func TestTaskIDFromString(t *testing.T) {
	t.Parallel()

	strUUID := "07cd2aba-8e5d-4827-ac43-66274ec4a2f4"
	expectedUUID := uuid.MustParse(strUUID)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id      string
		want    TaskID
		wantErr bool
	}{
		{
			name:    "successfully returns the TaskID ",
			id:      strUUID,
			want:    TaskID{value: expectedUUID},
			wantErr: false,
		},
		{
			name:    "rejects invaild UUID",
			id:      "INVALID",
			want:    TaskID{},
			wantErr: true,
		},
		{
			name:    "rejects empty string",
			id:      "",
			want:    TaskID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := TaskIDFromString(tt.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("TaskIDFromString() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("TaskIDFromString() succeeded unexpectedly")
			}

			if got.value != tt.want.value {
				t.Errorf("TaskIDFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
