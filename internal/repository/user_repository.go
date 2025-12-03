package repository

import (
	"errors"
	"gew/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(string) (*model.User, error)
	SaveUser(model.User) (*model.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUser(Db *gorm.DB) UserRepository {
	return &userRepository{
		Db,
	}
}

func (r *userRepository) FindUserByEmail(userEmail string) (*model.User, error) {
	var user model.User
	user.Email = userEmail
	err := r.Db.Where("email", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Record not found")
	}

	return &user, nil
}

func (r *userRepository) SaveUser(userInput model.User) (*model.User, error) {
	var user model.User
	user = userInput
	res := r.Db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
