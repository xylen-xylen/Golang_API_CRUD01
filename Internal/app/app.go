package app

import (
	"Golang_API_CRUD01/Internal/api"
	"Golang_API_CRUD01/Internal/store"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger      *log.Logger
	PostHandler *api.PostHandler
	DB          *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	postgresDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	postStore := store.NewPostgresPostStore(postgresDB)

	postHandler := api.NewPostHandler(postStore)

	app := &Application{
		Logger:      logger,
		PostHandler: postHandler,
		DB:          postgresDB,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is ok.")

}
