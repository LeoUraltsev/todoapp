package logger

import (
	"log"

	"go.uber.org/zap"
)

var l *zap.Logger

type Logger struct {
	*zap.Logger
}

func GetLogger() *Logger {
	return &Logger{l}
}

func init() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	l = logger
}
