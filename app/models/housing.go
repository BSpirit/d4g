package models

import (
	"d4g/app/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
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

	var houses []map[string]string
	for rows.Next() {
		var id, streetNumber, streetName, cityPostalCode, cityName string
		err := rows.Scan(&id, &streetNumber, &streetName, &cityPostalCode, &cityName)
		if err != nil {
			return "", utils.Trace(err)
		}

		house := map[string]string{
			"id":             id,
			"streetNumber":   streetNumber,
			"streetName":     streetName,
			"cityPostalCode": cityPostalCode,
			"cityName":       cityName}

		houses = append(houses, house)
	}
	result, err := json.Marshal(houses)
	if err != nil {
		return "", utils.Trace(err)
	}
	return string(result), nil
}

func GetHousingDetails(pk string, db *sql.DB) (string, error) {
	rows, err := db.Query(`SELECT c.housing_id, c.power_kw, c.date,
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

	type rowDetails struct {
		housingId string
		powerKw   int
		date      time.Time

		tenantFirstname string
		tenantLastname  string

		landlordLastname  string
		landlordFirstname string
		company           string
		address           string

		streetNumber   string
		streetName     string
		cityPostalCode string
		cityName       string
	}
	var detailsResult []map[string]interface{}

	for rows.Next() {
		var details rowDetails
		err := rows.Scan(&details.housingId, &details.powerKw, &details.date,
			&details.tenantFirstname, &details.tenantLastname,
			&details.landlordFirstname, &details.landlordLastname, &details.company, &details.address,
			&details.streetNumber, &details.streetName, &details.cityPostalCode, &details.cityName)
		rowDetails := map[string]interface{}{"id": details.housingId, "KW": details.powerKw, "date": details.date,
			"tenantFirstName": details.tenantFirstname, "tenantLastName": details.tenantLastname,
			"landlordFirstName": details.landlordFirstname, "landlordLastName": details.landlordLastname,
			"company": details.company, "address": details.address,
			"streetNumber": details.streetNumber, "streetName": details.streetName,
			"cityPostalCode": details.cityPostalCode, "cityName": details.cityName}

		if err != nil {
			return "", utils.Trace(err)
		}
		detailsResult = append(detailsResult, rowDetails)
	}
	result, err := json.Marshal(detailsResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	return string(result), nil
}
