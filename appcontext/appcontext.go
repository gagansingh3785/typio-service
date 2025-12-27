package appcontext

import "github.com/gagansingh3785/typio-service/config"

var AppCtx *AppContext

type AppContext struct {
	Cfg *config.Config
}

func Initiate(cfg *config.Config) {
	AppCtx = &AppContext{
		Cfg: cfg,
	}
}

func GetConfig() *config.Config {
	return AppCtx.Cfg
}
