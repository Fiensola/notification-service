package main

import (
	"os"

	"github.com/fiensola/notification-service/internal/app"
	"github.com/fiensola/notification-service/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	server := app.GetServer(":8080")

	logger.SugaredLogger.Info("Starting HTTPS server")

	if err := server.ListenAndServe(); err != nil {
		logger.SugaredLogger.Errorf("Failed to start HTTP server: %v", err)
		os.Exit(1)
	}
}
