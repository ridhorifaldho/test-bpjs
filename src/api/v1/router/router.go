package router

import (
	"github.com/gin-gonic/gin"
	"test-bpjs/config/db"
	"test-bpjs/config/env"
	ctrlRepo "test-bpjs/src/api/v1/controllers/order"
	repoOrder "test-bpjs/src/api/v1/repository/order"
	repoTrx "test-bpjs/src/api/v1/repository/transactional_db"
	uscOrder "test-bpjs/src/api/v1/usecase/order"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	DB := db.NewDB(env.Config).DB

	v1 := router.Group("v1.0")

	transactionalRepo := repoTrx.NewTransactionalDBRepoImpl(DB)
	orderRepo := repoOrder.NewRepoOrderImpl(DB)

	routerOrder := v1.Group("/order")
	orderUsc := uscOrder.OrderUsecaseImpl(orderRepo, transactionalRepo)
	ctrlRepo.OrderControllerImpl(routerOrder, orderUsc)

	return router
}
