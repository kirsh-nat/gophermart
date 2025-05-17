package main

import (
	"embed"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kirsh-nat/gophermart.git/gophermart/cmd/gophermart/migrations"
	_ "github.com/kirsh-nat/gophermart.git/gophermart/cmd/gophermart/migrations"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/app"
	"github.com/kirsh-nat/gophermart.git/gophermart/internal/handlers"
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
	return http.ListenAndServe("localhost:8080", mux)
}
