package models

import (
	"d4g/app/utils"
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
)

type Housing struct {
	HousingID     string `json:"id"`
	Type          int	`json:"type"`
	SurfaceArea   int	`json:"surfaceArea"`
	Rooms         int	`json:"roomsNb"`
	HeatingSystem string `json:"heatingSystem"`
	Year          int	`json:"constructionYear"`
	StreetNumber  string `json:"streetNumber"`
	Street        string `json:"streetName"`
	Postcode      string `json:"cityPostalCode"`
	City          string `json:"cityName"`
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

	type Details struct{
		Housing Housing
		Consumptions []Consumption
	}

	rows, err := db.Queryx(`SELECT h.housing_id, h.street_number, h.street, h.postcode, h.city,
	h.type, h.surface_area, h.rooms, h.heating_system, h.year,
	c.consumption_id, c.power_kw, c.date 
	FROM housing as h INNER JOIN consumption as c ON h.housing_id = c.housing_id 
	WHERE h.housing_id = ?`, pk)
	if err != nil {
		return "", utils.Trace(err)
	}
	defer rows.Close()

	var house Housing
	var consumptions []Consumption
	for rows.Next() {
		var conso Consumption
		err := rows.Scan(&house.HousingID, &house.StreetNumber, &house.Street, &house.Postcode, &house.City,
			&house.Type, &house.SurfaceArea, &house.Rooms, &house.HeatingSystem, &house.Year,
			&conso.ConsumptionID, &conso.PowerKW, &conso.Date)
		if err != nil {
			return "", utils.Trace(err)
		}
		consumptions = append(consumptions, Consumption{
			ConsumptionID: conso.ConsumptionID,
			HousingID:     house.HousingID,
			PowerKW:       conso.PowerKW,
			Date:          conso.Date,
		})
	}
	houseResult := Housing{
		HousingID: house.HousingID,
		Type: house.Type,
		SurfaceArea: house.SurfaceArea,
		Rooms: house.Rooms,
		HeatingSystem: house.HeatingSystem,
		Year: house.Year,
		StreetNumber: house.StreetNumber,
		Street: house.Street,
		Postcode: house.Postcode,
		City: house.City,
	}
	details := Details{
		Housing:      houseResult,
		Consumptions: consumptions,
	}

	result, err := json.Marshal(details)
	if err != nil {
		return "", utils.Trace(err)
	}

	return string(result), nil
}
