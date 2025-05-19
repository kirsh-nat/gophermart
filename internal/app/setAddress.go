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
		"a", os.Getenv("RUN_ADDRESS"),
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if adr := os.Getenv("RUN_ADDRESS"); adr != "" {
		Address = adr
	}
}
