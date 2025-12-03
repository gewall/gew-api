package domain

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type AuthResponse struct {
	Message string `json:"message"`
}

type AuthErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"message"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (render *AuthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (err *AuthErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, err.HTTPStatusCode)
	return nil
}

func NewAuthResponse(message string) *AuthResponse {
	return &AuthResponse{Message: message}
}

func AuthErrInvalidRequest(err error) *AuthErrorResponse {
	return &AuthErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
func (bind *LoginRequest) Bind(r *http.Request) error {
	if bind.Email == "" || bind.Password == "" {
		return errors.New("Missing required fields")
	}
	return nil
}
