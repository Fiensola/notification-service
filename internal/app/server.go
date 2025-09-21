package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServer(addr string) *http.Server {
	router := getRouter()

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
