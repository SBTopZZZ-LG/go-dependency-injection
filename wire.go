//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"gorm.io/gorm"
	"todo_app/cli"
	"todo_app/cli/create_todo_command"
	"todo_app/cli/delete_todo_command"
	"todo_app/cli/get_todo_command"
	"todo_app/cli/list_todos_command"
	"todo_app/cli/make_todo_important_command"
	"todo_app/cli/make_todo_not_important_command"
	"todo_app/cli/update_todo_command"
	"todo_app/config"
	"todo_app/logger"
	"todo_app/logger/console_logger"
	"todo_app/logger/loggers"
	"todo_app/logger/zap_logger"
	"todo_app/repositories/to_do_repository"
	"todo_app/services/to_do_service"
	"todo_app/utils/gorm_util"
	"todo_app/utils/sql_util"
	"todo_app/utils/zap_util"
)

func provideSQLDB(dbConf *config.DBConfig) (*sql.DB, func(), error) {
	sqlDB, err := sql_util.CreateConnection(dbConf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		err := sql_util.CloseConnection(sqlDB)
		if err != nil {
			panic(err)
		}
	}

	return sqlDB, cleanup, nil
}

func InitGormWithMySQLDriver(dbConf *config.DBConfig) (*gorm.DB, func(), error) {
	wire.Build(
		provideSQLDB,
		gorm_util.NewSilentGormInstanceWithMySQLDriver,
	)

	return nil, nil, nil
}

func InitTodoService(gormDB *gorm.DB) to_do_service.ITODOService {
	wire.Build(
		to_do_repository.New,
		wire.Bind(new(to_do_repository.ITODORepository), new(*to_do_repository.TODORepository)),

		to_do_service.New,
		wire.Bind(new(to_do_service.ITODOService), new(*to_do_service.TODOService)),
	)
	return nil
}

func provideLoggersWithZapAndConsoleLogger(
	zapLogger *zap_logger.ZapLogger,
	consoleLogger *console_logger.ConsoleLogger,
) logger.ILogger {
	return loggers.New(zapLogger, consoleLogger)
}

func InitCompositeLoggerWithZapAndConsole(loggerConf *config.LoggerConfig) (logger.ILogger, error) {
	wire.Build(
		zap_util.NewZapLogger,
		zap_logger.New,

		wire.Value(&console_logger.ConsoleLoggerParams{MethodNamespaceSkip: 3}),
		console_logger.New,

		provideLoggersWithZapAndConsoleLogger,
	)
	return nil, nil
}

func provideCLICommands(todoService to_do_service.ITODOService, logger logger.ILogger) []cli.ICommand {
	createToDoCmd := create_todo_command.New(todoService, logger)
	deleteToDoCmd := delete_todo_command.New(todoService, logger)
	getToDoCmd := get_todo_command.New(todoService, logger)
	listToDoCmd := list_todos_command.New(todoService, logger)
	updatedToDoCmd := update_todo_command.New(todoService, logger)
	makeToDoImportantCmd := make_todo_important_command.New(todoService, logger)
	makeToDoNotImportantCmd := make_todo_not_important_command.New(todoService, logger)
	cliCommands := []cli.ICommand{
		createToDoCmd,
		deleteToDoCmd,
		getToDoCmd,
		listToDoCmd,
		updatedToDoCmd,
		makeToDoImportantCmd,
		makeToDoNotImportantCmd,
	}

	return cliCommands
}

func InitCLICommands(todoService to_do_service.ITODOService, logger logger.ILogger) *cli.RootCommand {
	wire.Build(
		provideCLICommands,
		cli.NewRootCommand,
	)
	return nil
}
