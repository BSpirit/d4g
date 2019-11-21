package handlers

import (
	"d4g/app/models"
	"d4g/app/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)


func CreateConsumptionHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	power,_ := strconv.Atoi(r.FormValue("powerkw"))
	date, err := time.Parse("01/01/2000", r.FormValue("date"))

	if err != nil {
		fmt.Println(err)
	}
	consumption := models.Consumption{
		HousingID: r.FormValue("housingid"),
		PowerKW: power,
		Date: date,
	}

	db, err := env.DB.Begin()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}
	defer db.Rollback()

	exist := consumption.IsAlreadyExist(db)
	if exist != false {
		return &StatusError{Code: 409, Err: utils.Trace(fmt.Errorf("Already exists!"))}
	}

	err = consumption.Create(db)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	err = db.Commit()
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}