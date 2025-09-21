package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugaredLogger *zap.SugaredLogger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	SugaredLogger = logger.Sugar()
}

func Sync() {
	_ = SugaredLogger.Sync()
}
