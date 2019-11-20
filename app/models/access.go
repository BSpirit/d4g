package models

import (
	"d4g/app/utils"
	"database/sql"
)

type Access struct {
	AccessID  int64
	HousingID sql.NullString
	Login     string
	Password  string
	IsAdmin   bool
}

func (a *Access) Create(tx *sql.Tx) error {
	res, err := tx.Exec(`INSERT INTO access(login, password, housing_id, is_admin)
						 VALUES(?, ? ,?, ?)`, a.Login, a.Password, a.HousingID, a.IsAdmin)
	if err != nil {
		return utils.Trace(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return utils.Trace(err)
	}
	a.AccessID = id

	return nil
}
