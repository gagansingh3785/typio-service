package appcontext

import (
	"github.com/gagansingh3785/typio-service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SetupDatabase(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.GetDBURL())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
