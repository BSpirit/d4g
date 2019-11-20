package handlers

import (
	"d4g/app/utils"
	"net/http"
)

func IndexHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	err := env.Templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	return nil
}
