package app

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func setDBConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")

	str := "host=localhost port=" + port + " user=" + user + " password=" + pass + " dbname=" + name + " sslmode=disable"
	flag.StringVar(&ConnStr,
		"d", str,
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if conn := os.Getenv("DATABASE_URI"); conn != "" {
		ConnStr = conn
	}
}
