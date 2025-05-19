package main

import (
	"embed"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/cmd/gophermart/migrations"
	"github.com/kirsh-nat/gophermart.git/internal/app"
	"github.com/kirsh-nat/gophermart.git/internal/handlers"
)

var embedMigrations embed.FS

func main() {
	app.SetAppConfig()
	err := migrations.RunMigrations(app.ConnStr)
	if err != nil {
		panic(err)
	}
	//service := models.NewModel(app.DB)
	handler := handlers.NewURLHandler(app.DB)

	if err := run(handler); err != nil {
		app.Sugar.Fatalw(err.Error(), "event", "start server")
	}
	if app.DB != nil {
		defer app.DB.Close()

	}
}

func run(handler *handlers.URLHandler) error {
	mux := handlers.Routes(handler)
	return http.ListenAndServe(app.Address, mux)
}
