package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	defaultLog "log"
	"sync"
)

var (
	once   = new(sync.Once)
	Logger *zap.SugaredLogger
)

func SetupLogger() {
	once.Do(func() {
		log, err := newLogger()
		if err != nil {
			defaultLog.Fatalf("Can't initialize logger: %v", err)
		}
		Logger = log.Sugar()
		Logger.Info("New logger was created!")
	})
}

func newLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	log, err := config.Build()
	return log, err
}
