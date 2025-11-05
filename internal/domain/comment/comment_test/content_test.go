package comment_test

import (
	"testing"
	"todo/internal/domain/comment"
)

func TestContentVO(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid content",
			input:       "This is a comment",
			expectError: false,
		},
		{
			name:        "empty content",
			input:       "",
			expectError: true,
		},
		{
			name:        "spaces only",
			input:       "     ",
			expectError: true,
		},
		{
			name:        "normal content with punctuation",
			input:       "Hello, world!",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := comment.NewContent(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			if c.String() != tt.input {
				t.Errorf("expected content %q, got %q", tt.input, c.String())
			}
		})
	}
}
