package app

import (
	conf "EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/pkg/logging"
)

func Run() {

	logger := logging.GetLogger()
	logger.Info("Start application")
	config := conf.GetConfig()
	config.Storage.Host = "dasd"
	repositores := repos.NewRepositories(crud.ConnPool)

}
