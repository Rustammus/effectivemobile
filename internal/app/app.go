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
)

func Run() {
	logger := logging.GetLogger()
	logger.Info("Start application")
	conf := config.GetConfig()

	migration.DoMigrate()

	repositories := repos.NewRepositories(crud.GetPool())
	allService := service.NewServices(service.Deps{Repos: repositories})

	server := gin.Default()
	r := route.NewHandler(allService)
	r.Init(server)

	err := server.Run(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		fmt.Println(err)
		panic("start app failure")
	}
}
