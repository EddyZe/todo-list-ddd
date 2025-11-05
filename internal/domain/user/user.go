package user

import "todo/internal/domain/common"

type User struct {
	id       common.ID
	email    Email
	password Password
}

func NewUser(email Email, password Password) *User {
	id := common.GenerateID()

	return &User{
		id:       id,
		email:    email,
		password: password,
	}
}

func (u *User) ChangeEmail(email Email) {
	u.email = email
}

func (u *User) ChangePassword(password Password) {
	u.password = password
}

func (u *User) ComparePassword(p string) bool {
	return u.password.Compare(p)
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Password() Password {
	return u.password
}

func (u *User) ID() common.ID {
	return u.id
}
