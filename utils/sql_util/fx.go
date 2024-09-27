package sql_util

import (
	"context"
	"database/sql"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"todo_app/logger"
)

func RegisterHooks(lc fx.Lifecycle, logger logger.ILogger, sqlDB *sql.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("attempting to ping the database")

			err := sqlDB.PingContext(ctx)
			if err != nil {
				logger.Error("failed to ping the database", zap.Error(err))
			}

			return err
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("attempting to close the database")

			err := CloseConnection(sqlDB)
			if err != nil {
				logger.Error("failed to close the database", zap.Error(err))
			}

			return err
		},
	})
}

var Module = fx.Module(
	"sql_util",
	fx.Provide(CreateConnection),
	fx.Invoke(RegisterHooks),
)
