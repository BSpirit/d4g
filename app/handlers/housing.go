package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"net/http"
)

func HousingHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	res, err := models.GetHousing(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(res))

	return nil
}

func DetailsHousingHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	id := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	res, err := models.GetHousingDetails(id, limit, env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(res))

	return nil
}

func AllDetailsHousingHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	res, err := models.GetAllHousingDetails(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(res))

	return nil
}
