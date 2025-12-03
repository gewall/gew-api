package service

import (
	"errors"
	"gew/internal/http/dto"
	"gew/internal/model"
	"gew/internal/repository"
)

type LinkService interface {
	CreateLink(*dto.LinkRequest) error
	FindBySlug(string) (map[string]any, error)
	FindsById(string) ([]*dto.LinkResponse, error)
	DeleteById(string) error
}

type linkService struct {
	repo repository.LinkRepository
}

func NewLink(repo repository.LinkRepository) LinkService {
	return &linkService{repo}
}

func (s *linkService) FindBySlug(slug string) (map[string]any, error) {
	res, err := s.repo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"destination": res.Destination,
	}, nil
}

func (s *linkService) CreateLink(linkInput *dto.LinkRequest) error {

	link := &model.Link{
		Slug:        linkInput.Slug,
		UserID:      linkInput.UserID,
		Destination: linkInput.Destination,
		Title:       linkInput.Title,
	}
	_, err := s.repo.SaveLink(*link)
	if err != nil {
		return err
	}

	return nil
}

func (s *linkService) FindsById(id string) ([]*dto.LinkResponse, error) {
	var links []*dto.LinkResponse
	if id == "" {
		return nil, errors.New("id is invlaid")
	}

	res, err := s.repo.FindsById(id)
	if err != nil {
		return nil, err
	}

	for _, link := range *res {
		cvtLink := dto.ConvertModelToDTOLinkResponse(link)
		links = append(links, cvtLink)
	}

	return links, nil

}

func (s *linkService) DeleteById(id string) error {
	if err := s.repo.DeleteById(id); err != nil {
		return err
	}

	return nil
}
