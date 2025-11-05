package user_test

import (
	"testing"
	"todo/internal/domain/user"
)

func TestPasswordVO(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "normal password",
			input:       "MySecret123!",
			expectError: false,
		},
		{
			name:        "password with spaces",
			input:       "   password123   ",
			expectError: false,
		},
		{
			name:        "empty password",
			input:       "",
			expectError: true,
		},
		{
			name:        "password with symbols",
			input:       "!@#$%^&*()_+",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw, err := user.NewPassword(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.input, err)
			}

			// Проверяем, что Compare возвращает true для оригинального пароля
			if !pw.Compare(tt.input) {
				t.Errorf("password %q did not match its hash", tt.input)
			}

			// Проверяем, что Compare возвращает false для неправильного пароля
			if pw.Compare(tt.input + "wrong") {
				t.Errorf("password %q incorrectly matched with wrong password", tt.input)
			}
		})
	}
}
