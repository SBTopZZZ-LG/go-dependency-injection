package loggers

import "todo_app/logger"

type Loggers struct {
	loggers []logger.ILogger

	logger.ILogger
}

func New(logger ...logger.ILogger) *Loggers {
	return &Loggers{
		loggers: logger,
	}
}

func (l *Loggers) Info(msg string, args ...interface{}) {
	for _, l := range l.loggers {
		l.Info(msg, args...)
	}
}

func (l *Loggers) Error(msg string, args ...interface{}) {
	for _, l := range l.loggers {
		l.Error(msg, args...)
	}
}

func (l *Loggers) Fatal(msg string, args ...interface{}) {
	for _, l := range l.loggers {
		l.Fatal(msg, args...)
	}
}
