package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mnctest.com/api/authenticator"
	"mnctest.com/api/config"
	"mnctest.com/api/delivery/api"
	"mnctest.com/api/delivery/middleware"
	"mnctest.com/api/manager"
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
	api.NewLoginApi(apiLogin, s.UseCaseManager.LoginCustomerUseCase(), s.ConfigToken)

	apiCustomer := s.gin.Group("customers")
	api.NewCustomerApi(apiCustomer, s.UseCaseManager.CustomerUseCase())
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
