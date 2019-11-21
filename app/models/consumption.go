package models

import (
	"d4g/app/utils"
	"database/sql"
	"time"
)

type Consumption struct {
	ConsumptionID int64 `json:"id"`
	HousingID     string `json:"housingId"`
	PowerKW       int `json:"powerKw"`
	Date          time.Time `json:"date"`
}

func (c *Consumption) Create(tx *sql.Tx) error {
	res, err := tx.Exec(`INSERT INTO consumption(housing_id, power_kw, date)
					     VALUES(?, ? ,?)`, c.HousingID, c.PowerKW, c.Date)
	if err != nil {
		return utils.Trace(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return utils.Trace(err)
	}
	c.ConsumptionID = id

	return nil
}
