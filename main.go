package main

import (
	"d4g/app/handlers"
	"d4g/app/models"
	"log"
	"net/http"
)

func main() {
	env := &handlers.Env{}
	var err error

	env.DB, err = models.InitDB("db/d4g.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}
	defer env.DB.Close()

	http.Handle("/csv", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.CSVHandler,
	})

	http.Handle("/housing", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.HousingHandler,
	})

	http.Handle("/detailsHousing", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.DetailsHousingHandler,
	})


	log.Fatal(http.ListenAndServe(":8080", nil))
}
