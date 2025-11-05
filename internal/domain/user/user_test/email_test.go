package user_test

import (
	"testing"
	"todo/internal/domain/user"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "valid email",
			input:       "test@example.com",
			expectError: false,
		},
		{
			name:        "empty email",
			input:       "",
			expectError: true,
		},
		{
			name:        "missing at symbol",
			input:       "testexample.com",
			expectError: true,
		},
		{
			name:        "missing domain",
			input:       "test@.com",
			expectError: true,
		},
		{
			name:        "invalid characters",
			input:       "test!@example.com",
			expectError: true,
		},
		{
			name:        "subdomain valid",
			input:       "user@mail.example.co.uk",
			expectError: false,
		},
		{
			name:        "dash in domain",
			input:       "user@my-domain.com",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := user.NewEmail(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			if string(email) != tt.input {
				t.Errorf("expected email %q, got %q", tt.input, email)
			}
		})
	}
}
