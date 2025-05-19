package app

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func setAddress() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to get .env fail")
	}
	flag.StringVar(&Address,
		"a", "localhost:8080",
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if Address == "" {
		Address = os.Getenv("RUN_ADDRESS")
	}

	Sugar.Info("Server address: ", Address)
}
