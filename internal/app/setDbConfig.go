package app

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func setDBConfig() {
	err := godotenv.Load("/opt/projects/gophermart/gophermart/cmd/gophermart/.env")
	if err != nil {
		log.Fatal("Fail to get .env fail")
	}
	flag.StringVar(&ConnStr,
		"d", os.Getenv("DATABASE_URI"),
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if conn := os.Getenv("DATABASE_URI"); conn != "" {
		ConnStr = conn
	}
}
