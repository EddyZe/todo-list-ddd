package task

import (
	"testing"
	"todo/internal/domain/task"
)

func TestDescriptionVO(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid description", "This is a task description", false},
		{"empty description", "", true},
		{"spaces only", "   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			desc, err := task.NewDescription(tt.input)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			if desc.String() != tt.input {
				t.Errorf("expected description %q, got %q", tt.input, desc.String())
			}
		})
	}
}
