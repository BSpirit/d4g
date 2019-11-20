package main

import (
	"d4g/app/handlers"
	"d4g/app/models"
	"html/template"
	"log"
	"net/http"
)

func main() {
	env := &handlers.Env{}
	var err error

	env.Templates = template.Must(template.ParseGlob("./templates/*/*.html"))
	env.DB, err = models.InitDB("db/d4g.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}
	defer env.DB.Close()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/housing_csv", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.HousingCSVHandler,
	})

	http.Handle("/access_csv", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.AccessCSVHandler,
	})

	http.Handle("/tenant_csv", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.TenantCSVHandler,
	})

	http.Handle("/landlord_csv", handlers.Handler{
		Env:         env,
		HandlerFunc: handlers.LandlordCSVHandler,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
