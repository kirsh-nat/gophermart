package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func setDbConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to get .env fail")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
}
