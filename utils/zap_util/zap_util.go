package zap_util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"todo_app/config"
)

func NewZapLogger(conf *config.LoggerConfig) (*zap.Logger, error) {
	logLevel, err := zap.ParseAtomicLevel(conf.Level)
	if err != nil {
		return nil, err
	}

	zapConf := zap.Config{
		Level:            logLevel,
		Development:      conf.Development,
		Encoding:         conf.Encoding,
		OutputPaths:      conf.OutputPaths,
		ErrorOutputPaths: conf.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     conf.EncoderConfig.LineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
	}

	logger, err := zapConf.Build()

	return logger, err
}
