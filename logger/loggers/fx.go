package loggers

import (
	"go.uber.org/fx"
	"todo_app/logger"
)

var Module = fx.Module(
	"loggers",
	fx.Provide(
		fx.Annotate(
			New,
			fx.As(new(logger.ILogger)),
			fx.ParamTags(`group:"loggers"`),
		),
	),
)
