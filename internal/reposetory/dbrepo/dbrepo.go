package dbrepo

import (
	"database/sql"

	"github.com/conrad3rd/bookings/internal/config"
	"github.com/conrad3rd/bookings/internal/reposetory"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) reposetory.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
