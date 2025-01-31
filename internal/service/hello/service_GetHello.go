package hello

import "github.com/zingerone/base_go/internal/repository/user"

func (s *HelloService) GetHello() GetHelloOutput {
	userDetail, _ := s.userRepository.GetUserDetail(user.GetUserDetailInput{})
	return GetHelloOutput{
		Id:    userDetail.UserID,
		Name:  userDetail.Username,
		Email: userDetail.Email,
	}
}
