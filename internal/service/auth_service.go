package service

import "gew/internal/repository"

type AuthService interface {
	Ping() error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuth(repo repository.UserRepository) AuthService {
	return &authService{
		repo,
	}
}

func (s *authService) Ping() error {
	return nil
}
