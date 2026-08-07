package task

import "testing"

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

			if got.String() != tt.want {
				t.Errorf("error while creating taskTitle, got %v ,want %v", got.value, tt.want)
			}
		})
	}
}
