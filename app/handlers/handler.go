package handlers

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se *StatusError) Error() string {
	return se.Err.Error()
}

func (se *StatusError) Unwrap() error {
	return se.Err
}

type Env struct {
	DB     *sqlx.DB
	JWTKey []byte
}

type Handler struct {
	*Env
	HandlerFunc func(e *Env, w http.ResponseWriter, r *http.Request) *StatusError
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.HandlerFunc(h.Env, w, r)
	if err != nil {
		log.Printf("%s", err.Unwrap())
		http.Error(w, http.StatusText(err.Code), err.Code)
	}
}
