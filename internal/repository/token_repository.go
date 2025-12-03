package repository

import (
	"gew/internal/model"

	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(*model.RefreshToken) error
	FindToken(string) (*model.RefreshToken, error)
	UpdateToken(string) error
}

type tokenRepository struct {
	Db *gorm.DB
}

func NewToken(Db *gorm.DB) TokenRepository {
	return &tokenRepository{
		Db,
	}
}

func (r *tokenRepository) SaveToken(token *model.RefreshToken) error {
	var tokenData model.RefreshToken
	tokenData = *token
	if err := r.Db.Create(&tokenData).Error; err != nil {
		return err
	}

	return nil
}

func (r *tokenRepository) FindToken(refrToken string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	if err := r.Db.Where("token = ? AND revoked = false", refrToken).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) UpdateToken(refrToken string) error {
	var token model.RefreshToken
	if err := r.Db.Model(&token).Where("token = ?", refrToken).Update("revoked", true).Error; err != nil {
		return err
	}
	return nil
}
