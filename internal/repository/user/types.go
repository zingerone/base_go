package user

type GetUserDetailInput struct {
	UserID int
}

type GetUserDetailOutput struct {
	UserID   string
	Username string
	Email    string
}
