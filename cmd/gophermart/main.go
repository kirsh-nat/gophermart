package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/cmd/gophermart/migrations"
	"github.com/kirsh-nat/gophermart.git/internal/app"
	"github.com/kirsh-nat/gophermart.git/internal/handlers"
)

var embedMigrations embed.FS

func main() {
	app.SetAppConfig()
	app.Sugar.Info("Start server")
	err := migrations.RunMigrations(app.ConnStr)
	if err != nil {
		app.Sugar.Fatalw(err.Error(), "event", "start db")

	}
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
	fmt.Println("Server is running on", app.Address)
	return http.ListenAndServe(app.Address, mux)
}
