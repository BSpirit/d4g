package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"encoding/csv"
	"net/http"
	"os"
)

func LandlordCSVHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	csvFile, err := os.Open("db/data/landlord.csv")
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	tx, err := env.DB.Begin()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}
	defer tx.Rollback()

	for i, record := range records {
		if i == 0 {
			continue
		}
		
		landlord := &models.Landlord{
			HousingID:     record[0],
			Firstname:     record[2],
			Lastname:	   record[1],
			Company:	   record[3],
			Address:	   record[4],
		}

		err = landlord.Create(tx)
		if err != nil {
			return &StatusError{Code: 500, Err: utils.Trace(err)}
		}

	}

	err = tx.Commit()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	return nil
}
