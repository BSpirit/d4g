package models

import (
	"d4g/app/utils"
	"database/sql"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type Housing struct {
	HousingID     string
	Type          int
	SurfaceArea   int
	Rooms         int
	HeatingSystem string
	Year          int
	StreetNumber  string
	Street        string
	Postcode      string
	City          string
}

func (h *Housing) Create(tx *sql.Tx) error {
	_, err := tx.Exec(`INSERT INTO housing(housing_id, type, surface_area,
											 rooms, heating_system, year,
											 street_number, street,
											 postcode, city)
											 VALUES(?, ? ,?, ? ,?, ? ,?, ? ,?, ?)`,
		h.HousingID, h.Type, h.SurfaceArea,
		h.Rooms, h.HeatingSystem, h.Year,
		h.StreetNumber, h.Street,
		h.Postcode, h.City)

	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func GetHousing(db *sqlx.DB) (string, error) {
	rows, err := db.Queryx("SELECT housing_id, street_number, street, postcode, city FROM housing")
	if err != nil {
		return "", utils.Trace(err)
	}
	defer rows.Close()

	var houses []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			return "", utils.Trace(err)
		}

		houses = append(houses, row)
	}
	result, err := json.Marshal(houses)
	if err != nil {
		return "", utils.Trace(err)
	}
	return string(result), nil
}

func GetHousingDetails(pk string, db *sqlx.DB) (string, error) {
	rows, err := db.Queryx(`SELECT c.housing_id, c.power_kw, c.date,
										t.firstname, t.lastname,
										l.lastname, l.firstname, l.company, l.address,
										h.street_number, h.street, h.postcode, h.city
								FROM consumption as c INNER JOIN tenant as t ON c.housing_id = t.housing_id 
								INNER JOIN landlord as l ON c.housing_id = l.housing_id 
								INNER JOIN housing as h ON h.housing_id = c.housing_id 
								WHERE c.housing_id  = ?`, pk)
	if err != nil {
		return "", utils.Trace(err)
	}
	defer rows.Close()

	var detailsResult []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			return "", utils.Trace(err)
		}

		detailsResult = append(detailsResult, row)
	}
	result, err := json.Marshal(detailsResult)
	if err != nil {
		return "", utils.Trace(err)
	}

	return string(result), nil
}
