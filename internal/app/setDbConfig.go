package app

import (
	"flag"
	"os"
)

func setDBConfig() {
	//err := godotenv.Load("/opt/projects/gophermart/gophermart/cmd/gophermart/.env")
	// if err != nil {
	// 	log.Fatal("Fail to get .env fail")
	// }
	flag.StringVar(&ConnStr,
		"d", "host=localhost port=5432 user=gophermart password=password123 dbname=gophermart sslmode=disable",
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if conn := os.Getenv("DATABASE_URI"); conn != "" {
		ConnStr = conn
	}
}
