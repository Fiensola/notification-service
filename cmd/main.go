package main

import (
	"os"

	"github.com/Fiensola/notification-service/internal/app"
	"github.com/Fiensola/notification-service/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	if err := app.StartServer(":8080"); err != nil {
		logger.SugaredLogger.Errorf("Failed to start HTTP server: %v", err)
		os.Exit(1)
	}
}