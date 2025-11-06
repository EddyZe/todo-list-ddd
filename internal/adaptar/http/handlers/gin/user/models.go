package user

type AllUserResp struct {
	Users       []User `json:"users"`
	TotalPages  int    `json:"total_pages"`
	Limit       int    `json:"limit"`
	CurrentPage int    `json:"current_page"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
