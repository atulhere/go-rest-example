package main

import (
	"log"
	"net/http"

	"github.com/atulhere/go-rest-example/app"
	"github.com/atulhere/go-rest-example/models"
	"github.com/gorilla/mux"
)

func main() {

	database, err := models.ConnectDatabase()

	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
