package hello

import "github.com/zingerone/base_go/internal/repository/user"

type HelloService struct {
	userRepository user.UserRepositoryInterface
}

func NewHelloService() *HelloService {
	return &HelloService{
		userRepository: user.NewUserRepository(),
	}
}
