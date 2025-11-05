package user

type InputUser struct {
	Email    string
	Password string
}

type OutputUser struct {
	ID    string
	Email string
}

type InputUserList struct {
	Page  int
	Limit int
}

type OutputUserList struct {
	Users       []*OutputUser
	CurrentPage int
	Limit       int
	TotalPages  int
}
