package zap_logger

import (
	"go.uber.org/zap"
	"todo_app/logger"
)

type ZapLogger struct {
	logger *zap.Logger

	logger.ILogger
}

func New(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (zl *ZapLogger) Info(msg string, args ...interface{}) {
	var fields []zap.Field
	for _, arg := range args {
		fields = append(fields, arg.(zap.Field))
	}

	zl.logger.Info(msg, fields...)
}

func (zl *ZapLogger) Error(msg string, args ...interface{}) {
	var fields []zap.Field
	for _, arg := range args {
		fields = append(fields, arg.(zap.Field))
	}

	zl.logger.Error(msg, fields...)
}

func (zl *ZapLogger) Fatal(msg string, args ...interface{}) {
	var fields []zap.Field
	for _, arg := range args {
		fields = append(fields, arg.(zap.Field))
	}

	zl.logger.Fatal(msg, fields...)
}
