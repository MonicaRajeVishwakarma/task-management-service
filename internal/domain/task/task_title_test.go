package task

import (
	"testing"
)

func TestNewTaskTitle(t *testing.T) {
	t.Parallel()

	validTitle := "valid title"

	tests := []struct {
		name    string
		title   string
		want    string
		wantErr bool
	}{
		{
			name:    "creates title from valid input",
			title:   validTitle,
			want:    validTitle,
			wantErr: false,
		},
		{
			name:    "rejects empty title",
			title:   "",
			wantErr: true,
		},
		{
			name:    "rejects whitespace title",
			title:   "      ",
			wantErr: true,
		},
		{
			name:    "trim the surrounding whitespace",
			title:   "    " + validTitle + "   ",
			want:    validTitle,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTaskTitle(tt.title)

			if (err != nil) != tt.wantErr {
				t.Fatal("error while creating taskTitle")
			}

			if got.value != tt.want {
				t.Errorf("error while creating taskTitle, got %v ,want %v", got.value, tt.want)
			}
		})
	}
}

// TestTaskTitle_String provides the String value of the TestTitle
func TestTaskTitle_String(t *testing.T) {
	t.Parallel()

	inputTitle := TaskTitle{
		value: "Title",
	}
	expectedTitle := "Title"

	if inputTitle.String() != expectedTitle {
		t.Errorf("inputTitle %v is not matching with expected title %v ", inputTitle, expectedTitle)
	}
}

// TestTaskTitle_String compares the titles and return equality results
func TestTaskTitle_Equals(t *testing.T){
	t.Parallel()

	title := TaskTitle{
		value: "Title",
	}

	otherTitle := TaskTitle{
		value: "Title",
	}

	if !title.Equals(otherTitle){
		t.Errorf("title %v and otherTitle %v are not equal",title,otherTitle)
	}
}
