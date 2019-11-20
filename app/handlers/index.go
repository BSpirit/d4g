package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"net/http"
)

func CSVHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	err := models.InsertHousings(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = models.InsertAccess(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = models.InsertLandlord(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = models.InsertTenant(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = models.InsertConsumption(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	return nil
}
