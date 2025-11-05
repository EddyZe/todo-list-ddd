package passencode_test

import (
	"testing"
	"todo/internal/pkg/passencode"
)

func TestHashAndComparePassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectCompare bool
	}{
		{
			name:          "normal password",
			password:      "MySecret123!",
			expectCompare: true,
		},
		{
			name:          "empty password",
			password:      "",
			expectCompare: true,
		},
		{
			name:          "password with symbols",
			password:      "!@#$%^&*()_+",
			expectCompare: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			t.Log(tt.password)

			ok, hashed := passencode.HashPassword(tt.password)
			if !ok {
				t.Fatalf("failed to hash password %q", tt.password)
			}

			t.Log(hashed)

			// сравниваем оригинальный пароль с хэшем
			if !passencode.ComparePasswords(hashed, tt.password) {
				t.Errorf("password %q did not match its hash", tt.password)
			}

			// проверяем, что не совпадает с другим паролем
			if passencode.ComparePasswords(hashed, tt.password+"wrong") {
				t.Errorf("password %q incorrectly matched with wrong password", tt.password)
			}
		})
	}
}
