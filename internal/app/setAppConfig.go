package app

import (
	"database/sql"
)

func SetAppConfig() {
	setLogger()
	setDbConfig()
	var err error
	//storage := Storage{}
	DB, err = sql.Open("pgx", ConnStr)
	if err != nil {
		panic(err)
	}
}
