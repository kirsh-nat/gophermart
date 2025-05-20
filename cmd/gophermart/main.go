package main

import (
	"context"
	"embed"
	"net/http"
	"time"

	"github.com/kirsh-nat/gophermart.git/cmd/gophermart/migrations"
	"github.com/kirsh-nat/gophermart.git/internal/app"
	"github.com/kirsh-nat/gophermart.git/internal/handlers"
	"github.com/kirsh-nat/gophermart.git/internal/models/order"
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

	go func() {
		for {
			order.Worker(app.DB, context.Background(), app.AccrualAddress)
			time.Sleep(3 * time.Second)
		}
	}()

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
