package user_test

import (
	"testing"
	"todo/internal/domain/user"
)

func TestUserAggregate(t *testing.T) {
	t.Run("create user and check fields", func(t *testing.T) {
		email, _ := user.NewEmail("test@example.com")
		password, _ := user.NewPassword("MySecret123!")

		u := user.NewUser(email, password)

		if u.Email() != email {
			t.Errorf("expected email %q, got %q", email, u.Email())
		}

		if !u.ComparePassword("MySecret123!") {
			t.Errorf("expected password to match")
		}

		if u.ComparePassword("wrong") {
			t.Errorf("expected password not to match")
		}
	})

	t.Run("change email", func(t *testing.T) {
		email, _ := user.NewEmail("test@example.com")
		password, _ := user.NewPassword("MySecret123!")
		u := user.NewUser(email, password)

		newEmail, _ := user.NewEmail("new@example.com")
		u.ChangeEmail(newEmail)

		if u.Email() != newEmail {
			t.Errorf("expected email to be changed to %q, got %q", newEmail, u.Email())
		}
	})

	t.Run("change password", func(t *testing.T) {
		email, _ := user.NewEmail("test@example.com")
		password, _ := user.NewPassword("MySecret123!")
		u := user.NewUser(email, password)

		newPassword, _ := user.NewPassword("NewSecret456!")
		u.ChangePassword(newPassword)

		if !u.ComparePassword("NewSecret456!") {
			t.Errorf("expected new password to match")
		}

		if u.ComparePassword("MySecret123!") {
			t.Errorf("expected old password not to match")
		}
	})
}
