package main

import (
	"go.uber.org/fx"
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
	"todo_app/entities"
	"todo_app/logger/console_logger"
	"todo_app/logger/loggers"
	"todo_app/logger/zap_logger"
	"todo_app/repositories/to_do_repository"
	"todo_app/services/to_do_service"
	"todo_app/utils/gorm_util"
	"todo_app/utils/sql_util"
	"todo_app/utils/zap_util"
)

func main() {
	fx.New(
		//TIP Disable Fx Logging
		fx.NopLogger,

		//TIP <h5>App Config</h5>
		config.Module,

		//TIP <h5>Logging</h5>
		//TIP Zap Logger
		zap_util.Module,
		zap_logger.Module,
		//TIP Console Logger
		fx.Supply(&console_logger.ConsoleLoggerParams{
			MethodNamespaceSkip: 2,
		}),
		console_logger.Module,
		//TIP Composite Logger
		loggers.Module,

		//TIP <h5>DB & ORM</h5>
		sql_util.Module,
		gorm_util.Module,

		//TIP <h5>App Services</h5>
		to_do_repository.Module,
		to_do_service.Module,

		//TIP <h5>CLI Commands</h5>
		create_todo_command.Module,
		delete_todo_command.Module,
		get_todo_command.Module,
		list_todos_command.Module,
		make_todo_important_command.Module,
		make_todo_not_important_command.Module,
		update_todo_command.Module,
		cli.Module,

		//TIP <h5>Main</h5>
		fx.Invoke(
			func(gormDB *gorm.DB, rootCmd *cli.RootCommand, shutdown fx.Shutdowner) {
				//TIP <h5>Run GORM Migration</h5>
				// This will create/update the To-do table automatically
				err := gormDB.AutoMigrate(&entities.TODO{})
				if err != nil {
					panic(err)
				}

				//TIP <h4>Execute the CLI</h4>
				err = rootCmd.Execute()
				if err != nil {
					panic(err)
				}

				//TIP <h5>Gracefully exit the Fx server</h5>
				err = shutdown.Shutdown()
				if err != nil {
					panic(err)
				}
			},
		),
	).Run()
}
