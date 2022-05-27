package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Hello from zap.",
		// Structured context as strongly typed Field values.
		zap.String("logger", "zap"),
		zap.Duration("backoff", time.Second),
	)
}
