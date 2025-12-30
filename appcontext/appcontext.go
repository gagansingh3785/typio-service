package appcontext

import (
	"github.com/gagansingh3785/typio-service/config"
	"github.com/gagansingh3785/typio-service/database"
)

var AppCtx *AppContext

type AppContext struct {
	Cfg *config.Config
	DB  *database.Database
}

func Initiate(cfg *config.Config) error {
	// setup database
	db, err := database.NewDatabase(cfg)
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

func GetDatabase() *database.Database {
	return AppCtx.DB
}
