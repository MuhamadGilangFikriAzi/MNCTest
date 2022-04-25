package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gokost.com/m/authenticator"
	"gokost.com/m/config"
	"gokost.com/m/delivery/api"
	"gokost.com/m/delivery/middleware"
	"gokost.com/m/manager"
)

type AppServer interface {
	Run()
}

type serverConfig struct {
	gin            *gin.Engine
	Name           string
	Port           string
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	Config         *config.Config
	Middleware     *middleware.AuthTokenMiddleware
	ConfigToken    authenticator.Token
}

func (s *serverConfig) initHeader() {
	s.gin.Use(s.Middleware.TokenAuthMiddleware())
	s.routeGroupApi()
}

func (s *serverConfig) routeGroupApi() {
	apiLogin := s.gin.Group("login")
	api.NewLoginApi(apiLogin, s.UseCaseManager.LoginAdminUseCase(), s.ConfigToken)
}

func (s *serverConfig) Run() {
	s.initHeader()
	s.gin.Run(fmt.Sprintf("%s:%s", s.Name, s.Port))
}

func Server() AppServer {
	ginStart := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config.ConfigDatabase)
	repo := manager.NewRepoManager(infra.PostgreConn(), infra.MysqlConn())
	usecase := manager.NewUseCaseManager(repo)
	configToken := infra.ConfigToken(config.ConfigToken)
	middleware := middleware.NewAuthTokenMiddleware(configToken)
	return &serverConfig{
		ginStart,
		config.ConfigServer.Url,
		config.ConfigServer.Port,
		infra,
		repo,
		usecase,
		config,
		middleware,
		configToken,
	}
}
