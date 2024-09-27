package console_logger

import (
	"go.uber.org/fx"
	"todo_app/logger"
)

var Module = fx.Module(
	"console_logger",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(logger.ILogger)),
			fx.ResultTags(`group:"loggers"`),
		),
	),
)
