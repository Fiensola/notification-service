package app

import (
	"net/http"

	"github.com/Fiensola/notification-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

func StartServer(addr string) error {
	router := getRouter()

	logger.SugaredLogger.Info("Starting HTTPS server on ", addr)
	
	return http.ListenAndServe(addr, router)
}

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
