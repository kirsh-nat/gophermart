package app

import (
	"database/sql"
)

func SetAppConfig() {
	setLogger()
	setDBConfig()
	setAddress()
	setAcrAddress()
	var err error
	DB, err = sql.Open("pgx", ConnStr)
	if err != nil {
		panic(err)
	}
}
