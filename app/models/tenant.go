package models

import (
	"d4g/app/utils"
	"database/sql"
)

type Tenant struct {
	TenantID	int64
	HousingID   string
	Firstname	string
	Lastname	string
}

func (t *Tenant) Create(tx *sql.Tx) error {
	res, err := tx.Exec(`INSERT INTO tenant(housing_id, firstname, lastname)
											 VALUES(? ,?, ? )`,
		t.HousingID, t.Firstname, t.Lastname)

	if err != nil {
		return utils.Trace(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return utils.Trace(err)
	}
	t.TenantID = id

	return nil
}
