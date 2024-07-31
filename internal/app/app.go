package app

import (
	_ "EffectiveMobile/docs"
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/crud"
	"EffectiveMobile/internal/repos"
	"EffectiveMobile/internal/route"
	"EffectiveMobile/internal/service"
	"EffectiveMobile/migration"
	"EffectiveMobile/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	logger := logging.GetLogger()
	logger.Info("Start application")
	conf := config.GetConfig()

	loglvl, err := logrus.ParseLevel(conf.Server.LogLevel)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Level = loglvl

	migration.NewMigrator(logger).Up()

	repositories := repos.NewRepositories(logger, crud.GetPool())
	allService := service.NewServices(service.Deps{Repos: repositories, Logger: logger})

	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())
	r := route.NewHandler(allService)
	r.Init(server)

	err = server.Run(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		logger.Fatalf("start application error: %s", err.Error())
	}
}
