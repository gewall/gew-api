package service

import (
	"errors"
	"gew/internal/config"
	"gew/internal/http/dto"
	"gew/internal/model"
	"gew/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Ping() error
	Login(*dto.LoginRequest) (map[string]any, error)
	Register(*dto.RegisterRequest) (map[string]any, error)
	Refresh(string) (map[string]any, error)
}

type authService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewAuth(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) AuthService {
	return &authService{
		userRepo, tokenRepo,
	}
}

func (s *authService) Ping() error {
	return nil
}

func (s *authService) Login(userInput *dto.LoginRequest) (map[string]any, error) {
	res, err := s.userRepo.FindUserByEmail(userInput.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.PasswordHash), []byte(userInput.Password)); err != nil {
		return nil, err
	}

	// generate token
	accessToken, err := config.GenerateToken(res.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := config.GenerateRandomString()

	// save refresh token
	refrToken := &model.RefreshToken{
		Token:     refreshToken,
		UserID:    res.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 3),
	}
	if err := s.tokenRepo.SaveToken(refrToken); err != nil {
		return nil, err
	}

	return map[string]any{
		"accessToken":     accessToken,
		"refreshToken":    refreshToken,
		"refreshTokenExp": (time.Hour * 24 * 3).Seconds(),
	}, nil

}

func (s *authService) Register(userInput *dto.RegisterRequest) (map[string]any, error) {
	// check email availabilty
	_, err := s.userRepo.FindUserByEmail(userInput.Email)
	if err == nil {
		return nil, errors.New("Email has been registered")
	}

	// generate hash password
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), 12)
	userInput.Password = string(hashPass)
	// convert dto to model
	userCnv := dto.ConvertToModelUser(userInput)

	// save user data
	res, err := s.userRepo.SaveUser(userCnv)
	if err != nil {
		return nil, err
	}

	// generate token
	accessToken, err := config.GenerateToken(res.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, _ := config.GenerateRandomString()

	// save refresh token
	refrToken := &model.RefreshToken{
		Token:     refreshToken,
		UserID:    res.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 3),
	}
	if err := s.tokenRepo.SaveToken(refrToken); err != nil {
		return nil, err
	}

	return map[string]any{
		"accessToken":     accessToken,
		"refreshToken":    refreshToken,
		"refreshTokenExp": (time.Hour * 24 * 3).Seconds(),
	}, nil
}

func (s *authService) Refresh(refrToken string) (map[string]any, error) {
	res, err := s.tokenRepo.FindToken(refrToken)
	if err != nil {
		return nil, err
	}
	if err := s.tokenRepo.UpdateToken(refrToken); err != nil {
		return nil, err
	}

	accessToken, err := config.GenerateToken(res.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken, _ := config.GenerateRandomString()

	// save refresh token
	refToken := &model.RefreshToken{
		Token:     refreshToken,
		UserID:    res.UserID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 3),
	}
	if err := s.tokenRepo.SaveToken(refToken); err != nil {
		return nil, err
	}

	return map[string]any{
		"accessToken":     accessToken,
		"refreshToken":    refreshToken,
		"refreshTokenExp": (time.Hour * 24 * 3).Seconds(),
	}, nil
}
