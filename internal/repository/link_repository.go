package repository

import (
	"errors"
	"gew/internal/model"

	"gorm.io/gorm"
)

type LinkRepository interface {
	SaveLink(model.Link) (*model.Link, error)
	FindBySlug(string) (*model.Link, error)
	FindsById(string) (*[]model.Link, error)
	DeleteById(string) error
}

type linkRepository struct {
	Db *gorm.DB
}

func NewLink(Db *gorm.DB) LinkRepository {
	return &linkRepository{Db}
}
func (r *linkRepository) FindBySlug(slug string) (*model.Link, error) {
	var link model.Link
	err := r.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("slug=?", slug).First(&link).Error
		if err != nil {
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		err = tx.Model(&link).Update("visits", gorm.Expr("visits + ?", 1)).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *linkRepository) SaveLink(linkInput model.Link) (*model.Link, error) {
	var link model.Link
	link = linkInput
	if err := r.Db.Create(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *linkRepository) FindsById(id string) (*[]model.Link, error) {
	var link []model.Link

	if err := r.Db.Where("user_id=?", id).Find(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *linkRepository) DeleteById(id string) error {
	var link model.Link
	if err := r.Db.Where("id=?", id).Delete(&link).Error; err != nil {
		return err
	}

	return nil
}
