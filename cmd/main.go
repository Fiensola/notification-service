package main

import (
	"os"

	"github.com/fiensola/notification-service/internal/app"
	"github.com/fiensola/notification-service/internal/config"
	"github.com/fiensola/notification-service/pkg/logger"
)

func main() {
	// logger
	logger.InitLogger()
	defer logger.Sync()

	// env
	config, err := config.LoadConfig()
	if err != nil {
		logger.SugaredLogger.Warnf("Failed load config: %v", err)
	}

	if (config.HttpAddr == "") {
		logger.SugaredLogger.Error("HTTP_ADDR is not set. Check env")
		os.Exit(1)
	}

	server := app.GetServer(config.HttpAddr)

	logger.SugaredLogger.Info("Starting HTTPS server")

	if err := server.ListenAndServe(); err != nil {
		logger.SugaredLogger.Errorf("Failed to start HTTP server on %s: %v", config.HttpAddr, err)
		os.Exit(1)
	}
}
