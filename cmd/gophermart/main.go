package main

import (
	"context"
	"embed"
	"net/http"
	"time"

	"github.com/kirsh-nat/gophermart.git/cmd/gophermart/migrations"
	"github.com/kirsh-nat/gophermart.git/internal/app"
	"github.com/kirsh-nat/gophermart.git/internal/handlers"
	systemservices "github.com/kirsh-nat/gophermart.git/internal/services/systemServices"
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

	// cmd := exec.Command("../accrual/accrual_linux_amd64", "-a="+app.AccrualAddress)

	// if err := cmd.Start(); err != nil {
	// 	app.Sugar.Error("Error to start system accrual: %v", err)
	// }

	// defer func() {
	// 	if err := cmd.Process.Kill(); err != nil {
	// 		app.Sugar.Error("Not possible to kill accrual process: %v", err)
	// 	} else {
	// 		app.Sugar.Error("Service accrual stopped")
	// 	}
	// 	cmd.Wait()
	// }()

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	sig := <-sigs
	// 	app.Sugar.Error("Got signal %s, shutting down...", sig)
	// 	os.Exit(0)
	// }()

	go func() {
		for {
			systemservices.Worker(app.DB, context.Background(), app.AccrualAddress)
			time.Sleep(5 * time.Second)
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
