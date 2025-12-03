package dto

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type AuthErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"message"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  any    `json:"error,omitempty"`
}

// Response
func (render *Response) Render(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func NewResponse(Message string, Data any) *Response {
	return &Response{
		Message,
		Data,
	}
}

// Error Response
func (err *AuthErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, err.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) *AuthErrorResponse {
	var error map[string]string

	jsonErr := json.Unmarshal([]byte(err.Error()), &error)
	if jsonErr != nil {
		return &AuthErrorResponse{
			Err:            err,
			HTTPStatusCode: 400,
			StatusText:     "Invalid request.",
			ErrorText:      err.Error(),
		}
	}
	return &AuthErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      error,
	}
}
