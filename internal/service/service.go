package service

import "github.com/zingerone/base_go/internal/service/hello"

func NewService() *Service {
	helloService := hello.NewHelloService()
	return &Service{
		HelloService: helloService,
	}
}

type Service struct {
	HelloService hello.HelloServiceInterface
}
