package log

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	return
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Logger() *zap.Logger {
	return logger
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}
