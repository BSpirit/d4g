package models

import (
	"d4g/app/utils"
	"database/sql"
	"encoding/json"
	"fmt"
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


func GetHousing(db *sql.DB) (string, error) {
	rows, err := db.Query("SELECT housing_id, street_number, street, postcode, city FROM housing")
	if err != nil {
		return "", utils.Trace(err)
	}
	defer rows.Close()
	var houses []map[string]interface{}
	for rows.Next() {
		house := map[string]interface{}{
			"id": "",
			"streetNumber": "",
			"streetName": "",
			"cityPostalCode": "",
			"cityName": "",
		}

		housingId := ""
		streetNumber := ""
		street := ""
		postalcode := ""
		city := ""

		err := rows.Scan(&housingId, &streetNumber, &street, &postalcode, &city)
		house["id"] = housingId
		house["streetNumber"] = streetNumber
		house["streetName"] = street
		house["cityPostalCode"] = postalcode
		house["cityName"] = city

		if err != nil {
			return "", utils.Trace(err)
		}
		houses = append(houses, house)
	}
	result, err := json.Marshal(houses)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	return string(result), nil
}
