package app

import (
	"flag"
	"os"
)

func setAddress() {
	flag.StringVar(&Address,
		"a", "localhost:8000",
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if Address == "" {
		Address = os.Getenv("RUN_ADDRESS")
	}

	Sugar.Info("Server address: ", Address)
}
