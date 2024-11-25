package utils

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type MyError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewMyError(status int, message string) MyError {
	return MyError{Status: status, Message: message}
}

func (e MyError) Error() string {
	return fmt.Sprintf("my error: %s and status: %d", e.Message, e.Status)
}

func Render(t templ.Component, w http.ResponseWriter, r *http.Request) error {
	return t.Render(r.Context(), w)
}

// TODO: Show error page to the user
func MakeHandler(f MyHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error("HTTP handler error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}
