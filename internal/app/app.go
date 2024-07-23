package app

import (
	conf "EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/logging"
)

func Run() {

	logger := logging.GetLogger()
	logger.Info("Start application")
	config := conf.GetConfig()
	config.Storage.Host = "dasd"
}
