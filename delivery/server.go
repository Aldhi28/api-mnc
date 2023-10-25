package delivery

import (
	"fmt"
	"log"

	"login-go/config"
	"login-go/delivery/controller"
	"login-go/delivery/middleware"
	"login-go/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (a *appServer) initController() {
	a.engine.Use(middleware.LogRequestMiddleware(a.log))
	controller.NewProductController(a.engine, a.useCaseManager.ProductUseCase())
	controller.NewCustomerController(a.engine, a.useCaseManager.CustomerUseCase())
	controller.NewBillController(a.engine, a.useCaseManager.BillUseCase())
	controller.NewUserController(a.engine, a.useCaseManager.UserUseCase())
	controller.NewAuthController(a.engine, a.useCaseManager.AuthUseCase())

}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	// <--                 <--         <--         <--
	// CONFIGURASI -> MEMBUAT KONEKSI -> REPOSITORY -> USECASE -> CONTROLLER
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	infraManager, errConnect := manager.NewInfraManager(cfg)
	if errConnect != nil {
		log.Fatalln("Error Connection : ", errConnect.Error())
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &appServer{
		engine:         engine,
		useCaseManager: useCaseManager,
		host:           host,
		log:            logrus.New(),
	}

}
