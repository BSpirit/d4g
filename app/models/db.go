package models

import (
	"crypto/sha512"
	"d4g/app/utils"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InsertHousings(db *sql.DB) error {
	csvFile, err := os.Open("db/data/housings.csv")
	if err != nil {
		return utils.Trace(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return utils.Trace(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.Trace(err)
	}
	defer tx.Rollback()

	for i, record := range records {
		if i == 0 {
			continue
		}
		housingType, err := strconv.Atoi(record[1])
		if err != nil {
			return utils.Trace(err)
		}
		surfaceArea, err := strconv.Atoi(record[2])
		if err != nil {
			return utils.Trace(err)
		}
		rooms, err := strconv.Atoi(record[3])
		if err != nil {
			return utils.Trace(err)
		}
		year, err := strconv.Atoi(record[5])
		if err != nil {
			return utils.Trace(err)
		}

		housing := &Housing{
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
			return utils.Trace(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func InsertAccess(db *sql.DB) error {
	csvFile, err := os.Open("db/data/access.csv")
	if err != nil {
		return utils.Trace(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return utils.Trace(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.Trace(err)
	}
	defer tx.Rollback()

	for i, record := range records {
		if i == 0 {
			continue
		}

		sha512 := sha512.New()
		sha512.Write([]byte(record[2]))
		access := &Access{
			HousingID: sql.NullString{String: record[0], Valid: true},
			Login:     record[1],
			Password:  string(sha512.Sum(nil)),
		}

		err = access.Create(tx)
		if err != nil {
			return utils.Trace(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func InsertTenant(db *sql.DB) error {
	csvFile, err := os.Open("db/data/tenant.csv")
	if err != nil {
		return utils.Trace(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return utils.Trace(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.Trace(err)
	}
	defer tx.Rollback()

	for i, record := range records {
		if i == 0 {
			continue
		}

		tenant := &Tenant{
			HousingID: record[0],
			Firstname: record[1],
			Lastname:  record[2],
		}

		err = tenant.Create(tx)
		if err != nil {
			return utils.Trace(err)
		}

	}

	err = tx.Commit()
	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func InsertLandlord(db *sql.DB) error {
	csvFile, err := os.Open("db/data/landlord.csv")
	if err != nil {
		return utils.Trace(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return utils.Trace(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.Trace(err)
	}
	defer tx.Rollback()

	for i, record := range records {
		if i == 0 {
			continue
		}

		landlord := &Landlord{
			HousingID: record[0],
			Firstname: record[2],
			Lastname:  record[1],
			Company:   record[3],
			Address:   record[4],
		}

		err = landlord.Create(tx)
		if err != nil {
			return utils.Trace(err)
		}

	}

	err = tx.Commit()
	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func InsertConsumption(db *sql.DB) error {
	jsonFile, err := os.Open("db/data/data.json")
	if err != nil {
		return utils.Trace(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return utils.Trace(err)
	}

	var rows []map[string]interface{}
	err = json.Unmarshal(byteValue, &rows)
	if err != nil {
		return utils.Trace(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.Trace(err)
	}
	defer tx.Rollback()

	for _, row := range rows {
		housindID := row["foyer"].(string)
		consumptionData := row["consumptionData"].([]interface{})

		for _, object := range consumptionData {
			object := object.(map[string]interface{})

			powerKW := int(object["consumption"].(float64))
			const layout = "02/01/2006"
			date, err := time.Parse(layout, object["date"].(string))
			if err != nil {
				return utils.Trace(err)
			}

			consumption := &Consumption{
				HousingID: housindID,
				PowerKW:   powerKW,
				Date:      date,
			}

			err = consumption.Create(tx)
			if err != nil {
				return utils.Trace(err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.Trace(err)
	}

	return nil
}

func NewNullInt64(s string) sql.NullInt64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: n,
		Valid: true,
	}
}
