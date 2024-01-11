package service

import "sql_go/pkg/repository"

type Service struct {
	repo repository.RepositoryInterface
}

type ServiceInterface interface {
	Login()
	ForgotPassword()
	ChangePassword()
	GetCodeByUsername()
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{
		repo: repo,
	}
}
