package appcontext

import (
	"github.com/gagansingh3785/typio-service/config"
	"github.com/jmoiron/sqlx"
)

var AppCtx *AppContext

type AppContext struct {
	Cfg *config.Config
	DB  *sqlx.DB
}

func Initiate(cfg *config.Config) error {
	// setup database
	db, err := SetupDatabase(cfg)
	if err != nil {
		return err
	}

	AppCtx = &AppContext{
		Cfg: cfg,
		DB:  db,
	}

	return nil
}

func GetConfig() *config.Config {
	return AppCtx.Cfg
}
