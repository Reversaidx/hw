package logger

import (
	"net/http"

	"go.uber.org/zap"
)

type Logger struct { // TODO
	logs    zap.Logger
	handler http.Handler
}

func New(level string) (*Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"/dev/stdout"}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	log := Logger{
		logs: *logger,
	}
	logger.Info("initializing")
	return &log, nil
}

func (l Logger) Info(msg string) {
	l.logs.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logs.Error(msg)
}

func (l Logger) Warn(msg string) {
	l.logs.Warn(msg)
}
