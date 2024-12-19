package dbrepo

import (
	"database/sql"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/config"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: app,
		DB:  conn,
	}
}
