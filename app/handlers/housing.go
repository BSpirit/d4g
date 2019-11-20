package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"encoding/csv"
	"net/http"
	"os"
	"strconv"
)

func HousingCSVHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	csvFile, err := os.Open("db/data/logements.csv")
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
		housingType, err := strconv.Atoi(record[1])
		if err != nil {
			return &StatusError{Code: 500, Err: utils.Trace(err)}
		}
		surfaceArea, err := strconv.Atoi(record[2])
		if err != nil {
			return &StatusError{Code: 500, Err: utils.Trace(err)}
		}
		rooms, err := strconv.Atoi(record[3])
		if err != nil {
			return &StatusError{Code: 500, Err: utils.Trace(err)}
		}
		year, err := strconv.Atoi(record[5])
		if err != nil {
			return &StatusError{Code: 500, Err: utils.Trace(err)}
		}

		housing := &models.Housing{
			HousingID:     record[0],
			Type:          housingType,
			SurfaceArea:   surfaceArea,
			Rooms:         rooms,
			HeatingSystem: record[4],
			Year:          year,
			StreetNumber:  record[6],
			Street:        record[7],
			Postcode:      record[8],
			City:          record[9],
		}

		err = housing.Create(tx)
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


func HousingHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	err := r.ParseForm()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	res, err := models.GetHousing(env.DB)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(res))

	return nil
}