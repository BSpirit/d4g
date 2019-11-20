package models

import (
	"d4g/app/utils"
	"database/sql"
	"time"
)

type Consumption struct {
	ConsumptionID int64
	HousingID     string
	PowerKW       int
	Date          time.Time
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
