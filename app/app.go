package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("get").
		Path("/residents").
		HandlerFunc(app.getAllResidents)
}

func (app *App) getAllResidents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	log.Println("You fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}
