package app

import (
	"flag"
	"os"
)

func setAddress() {
	//	err := godotenv.Load("/opt/projects/gophermart/gophermart/cmd/gophermart/.env")
	// if err != nil {
	// 	log.Fatal("Fail to get .env fail")
	// }
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
