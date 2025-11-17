package repository

import "gorm.io/gorm"

type UserRepository interface {
	FindUserByEmail()
}

type userRepository struct {
	Db *gorm.DB
}

func NewUser(Db *gorm.DB) UserRepository {
	return &userRepository{
		Db,
	}
}

func (r *userRepository) FindUserByEmail() {

}
