package task_test

import (
	"testing"
	"todo/internal/domain/task"
)

func TestTitleVO(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid title", "Task 1", false},
		{"empty title", "", true},
		{"spaces only", "   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, err := task.NewTitle(tt.input)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			if title.String() != tt.input {
				t.Errorf("expected title %q, got %q", tt.input, title.String())
			}
		})
	}
}
