package dto

import (
	"errors"
	"gew/internal/config"

	"gew/internal/model"
	"net/http"
)

type LinkRequest struct {
	Slug        string `json:"slug" validate:"required,max=255"`
	UserID      uint   `json:"user_id" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Title       string `json:"title" validate:"required,max=255"`
}

type LinkResponse struct {
	Slug        string `json:"slug"`
	UserID      uint   `json:"user_id"`
	Destination string `json:"destination"`
	Title       string `json:"title"`
}

func (b *LinkRequest) Bind(r *http.Request) error {
	if err := config.ValidateStruct(b); err != nil {
		jsonErr := Format(err)
		return errors.New(string(jsonErr))
	}
	return nil
}

func ConvertModelToDTOLinkResponse(link model.Link) *LinkResponse {
	return &LinkResponse{
		Slug:        link.Slug,
		UserID:      link.UserID,
		Destination: link.Destination,
		Title:       link.Title,
	}
}
