package models

import (
	"d4g/app/utils"
	"database/sql"
)

type Landlord struct {
	LandlordID int64
	HousingID  string
	Lastname   string
	Firstname  string
	Company    string
	Address    string
}

func (l *Landlord) Create(tx *sql.Tx) error {
	res, err := tx.Exec(`INSERT INTO landlord(housing_id, lastname, firstname, company, address)
											 VALUES(?, ? ,?, ? ,?)`,
		l.HousingID, l.Lastname, l.Firstname, l.Company, l.Address)

	if err != nil {
		return utils.Trace(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return utils.Trace(err)
	}
	l.LandlordID = id

	return nil
}
