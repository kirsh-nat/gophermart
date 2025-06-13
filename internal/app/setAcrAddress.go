package app

import (
	"flag"
	"os"
)

func setAcrAddress() {
	flag.StringVar(&AccrualAddress,
		"r", "localhost:8080",
		"Адрес запуска HTTP-сервера",
	)
	flag.Parse()

	if AccrualAddress == "" {
		AccrualAddress = os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	}

	Sugar.Info("Server address: ", AccrualAddress)
}
