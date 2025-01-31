package user

type UserRepositoryInterface interface {
	GetUserDetail(input GetUserDetailInput) (output GetUserDetailOutput, err error)
}
