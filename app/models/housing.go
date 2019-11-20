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

	type rowHouse struct {
		id string
		streetNumber string
		streetName string
		cityPostalCode string
		cityName string
	}
	var houses []map[string]string

	for rows.Next() {
		var house rowHouse
		err := rows.Scan(&house.id, &house.streetNumber, &house.streetName, &house.cityPostalCode, &house.cityName)
		rowHouse := map[string]string{"id" : house.id, "streetNumber": house.streetNumber, "streetName": house.streetName,
			"cityPostalCode": house.cityPostalCode, "cityName" :house.cityName}

		if err != nil {
			return "", utils.Trace(err)
		}
		houses = append(houses, rowHouse)
	}
	result, err := json.Marshal(houses)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return string(result), nil
}

/*
func getHousingDetails(pk int64, db *sql.DB) (string, error) {

	err := db.QueryRow(`SELECT * FROM consumption as c INNER JOIN tenant as t ON c.housing_id = t.housing_id 
								INNER JOIN landlord as l ON c.housing_id = l.housing_id 
								INNER JOIN housing as h ON h.housing_id = c.housing_id 
								WHERE c.housing_id  = ?`, pk).Scan(&user.ID, &user.Username, &user.Age)
	if err != nil {
		return nil, utils.Trace(err)
	}
	return string(result), nil
}
*/