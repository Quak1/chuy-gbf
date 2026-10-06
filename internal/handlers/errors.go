package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Error      error `json:"-"`
	StatusCode int   `json:"-"`

	ErrorMsg  string `json:"status"`
	ErrorText string `json:"error"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRender(err error) render.Renderer {
	return &ErrResponse{
		Error:      err,
		StatusCode: 422,
		ErrorMsg:   "Error rendering response.",
		ErrorText:  err.Error(),
	}
}

func ErrorServer(err error) render.Renderer {
	return &ErrResponse{
		Error:      err,
		StatusCode: 500,
		ErrorText:  err.Error(),
	}
}

var ErrorNotFound = &ErrResponse{
	StatusCode: 404,
	ErrorMsg:   "Resource not found.",
}

func ErrorInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Error:      err,
		StatusCode: 400,
		ErrorMsg:   "Invalid request.",
		ErrorText:  err.Error(),
	}
}

func ErrorParse(key string) render.Renderer {
	return &ErrResponse{
		StatusCode: http.StatusInternalServerError,
		ErrorMsg:   "Failed to parse key: " + key,
	}
}

var ErrorForbidden = &ErrResponse{
	StatusCode: http.StatusForbidden,
	ErrorMsg:   "You don't have access to this resource",
}
