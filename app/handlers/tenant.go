package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"encoding/csv"
	"net/http"
	"os"
)

func TenantCSVHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	csvFile, err := os.Open("db/data/tenant.csv")
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
		
		tenant := &models.Tenant{
			HousingID:     record[0],
			Firstname:     record[1],
			Lastname:	   record[2],
		}

		err = tenant.Create(tx)
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
