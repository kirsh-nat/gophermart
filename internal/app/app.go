package app

import (
	"database/sql"

	"go.uber.org/zap"
)

// пока есть только логер и база данных
// TODO: натсройка выавода логов в файл
// TODO: сделать енв с базовыми настройками (лог файл и строка подключения к базе )
var (
	Sugar   zap.SugaredLogger
	DB      *sql.DB
	ConnStr string
	Address string
)
