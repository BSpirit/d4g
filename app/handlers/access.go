package handlers

import (
	"crypto/sha512"
	"d4g/app/models"
	"d4g/app/utils"
	"encoding/csv"
	"net/http"
	"os"
)

func AccessCSVHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	csvFile, err := os.Open("db/data/access.csv")
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

		sha512 := sha512.New()
		sha512.Write([]byte(record[2]))
		access := &models.Access{
			HousingID: record[0],
			Login:     record[1],
			Password:  string(sha512.Sum(nil)),
		}

		err = access.Create(tx)
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
