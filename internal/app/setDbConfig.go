package app

import (
	"flag"
	"os"
)

func setDBConfig() {
	flag.StringVar(&ConnStr,
		"d", "host=localhost port=5432 user=gophermart password=password123 dbname=gophermart sslmode=disable",
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if conn := os.Getenv("DATABASE_URI"); conn != "" {
		ConnStr = conn
	}
}
