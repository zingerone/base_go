package user

// GetUserDetail this is for getting user detail
func (r *UserRepository) GetUserDetail(input GetUserDetailInput) (output GetUserDetailOutput, err error) {
	return GetUserDetailOutput{
		UserID:   "c3e0dcbb-94b4-48a5-8d65-b9c49cbb5503",
		Username: "imam.magribi",
		Email:    "imammagribi@gmail.com",
	}, nil
}
