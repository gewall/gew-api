package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"gew/internal/config"
	"gew/internal/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type RegisterRequest struct {
	Name string `json:"name" validate:"required,max=255"`
	LoginRequest
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type RegisterResponse struct {
	AccessToken string `json:"accessToken"`
}

func (b *LoginRequest) Bind(r *http.Request) error {

	if err := config.ValidateStruct(b); err != nil {
		jsonErr := Format(err)
		return errors.New(string(jsonErr))
	}
	return nil
}

func (b *RegisterRequest) Bind(r *http.Request) error {
	if b == nil {
		return errors.New(`{"error":"Field is empty"}`)
	}
	if err := config.ValidateStruct(b); err != nil {
		jsonErr := Format(err)
		return errors.New(string(jsonErr))
	}
	return nil
}

func Format(err error) []byte {
	errMap := make(map[string]string)
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		field := strings.ToLower(e.Field())

		switch e.Tag() {
		case "required":
			errMap[field] = "field is required"
		case "email":
			errMap[field] = "invalid email format"
		case "min":
			errMap[field] = fmt.Sprintf("minimum length is %s", e.Param())
		default:
			errMap[field] = "invalid value"
		}
	}
	jsonErr, _ := json.Marshal(errMap)
	return jsonErr
}

func ConvertToModelUser(dto *RegisterRequest) model.User {
	return model.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: dto.Password,
	}
}
