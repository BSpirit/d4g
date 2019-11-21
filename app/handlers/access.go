package handlers

import (
	"crypto/sha512"
	"d4g/app/models"
	"d4g/app/utils"
	"database/sql"
	"fmt"
	"net/http"
)

func  StringToSQLNull(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func CreateAccessHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	sha512 := sha512.New()
	sha512.Write([]byte(r.FormValue("password")))

	access := models.Access{
		HousingID: StringToSQLNull(r.FormValue("housingid")),
		Login: r.FormValue("login"),
		Password: string(sha512.Sum(nil)),
	}

	db, err := env.DB.Begin()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}
	defer db.Rollback()

	exist := access.IsAlreadyExist(db)
	if exist != false {
		return &StatusError{Code: 409, Err: utils.Trace(fmt.Errorf("Already exists!"))}
	}

	err = access.Create(db)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = db.Commit()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}