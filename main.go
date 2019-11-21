package main

import (
	"d4g/app/handlers"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	env := &handlers.Env{}
	env.DB = sqlx.MustConnect("sqlite3", "db/d4g.db?_foreign_keys=on")
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

	http.Handle("/allDetailsHousing", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.AllDetailsHousingHandler,
	})

	http.Handle("/createAccess", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.CreateAccessHandler,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
