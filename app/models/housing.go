package models

import (
	"d4g/app/utils"
	"database/sql"
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
