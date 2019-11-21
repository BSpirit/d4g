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
	enableCors(&w)
	err := h.HandlerFunc(h.Env, w, r)
	if err != nil {
		log.Printf("%s", err.Unwrap())
		http.Error(w, http.StatusText(err.Code), err.Code)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
