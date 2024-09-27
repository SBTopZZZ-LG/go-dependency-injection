package main

import (
	"database/sql"
	"github.com/a631807682/zerofield"
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
	"todo_app/logger/loggers"
	"todo_app/logger/zap_logger"
	"todo_app/repositories/to_do_repository"
	"todo_app/services/to_do_service"
	"todo_app/utils/gorm_util"
	"todo_app/utils/sql_util"
	"todo_app/utils/zap_util"
)

func main() {
	//TIP <h5>Load Configuration from the `config.yaml` file</h5>
	conf, err := config.Load(config.DefaultConfigFileName)
	if err != nil {
		panic(err)
	}

	//TIP <h5>Create a MySQL database connection using the loaded configuration</h5>
	sqlDB, err := sql_util.CreateConnection(conf.DBConfig)
	if err != nil {
		panic(err)
	}
	//TIP Make sure to close the database connection when `main` function exits
	defer func(db *sql.DB) {
		err := sql_util.CloseConnection(db)
		if err != nil {
			panic(err)
		}
	}(sqlDB)

	//TIP <h5>Create a GORM instance with the MySQL database connection</h5>
	gormDB, err := gorm_util.NewSilentGormInstanceWithMySQLDriver(sqlDB)
	if err != nil {
		panic(err)
	}
	err = gormDB.Use(zerofield.NewPlugin())
	if err != nil {
		panic(err)
	}

	//TIP <h5>Run GORM Migration</h5>
	// This will create/update the To-do table automatically
	err = gormDB.AutoMigrate(&entities.TODO{})
	if err != nil {
		panic(err)
	}

	//TIP <h5>Create the To-do's corresponding repository and service</h5>
	todoRepository := to_do_repository.New(gormDB)
	todoService := to_do_service.New(todoRepository)

	//TIP <h5>Create the Loggers for logging</h5>
	// Create a composite logger that can log into multiple `ILogger` implementations
	zap, err := zap_util.NewZapLogger(conf.LoggerConfig)
	if err != nil {
		panic(err)
	}
	zapLogger := zap_logger.New(zap)
	//consoleLogger := console_logger.New(2)
	compositeLoggers := loggers.New(
		zapLogger,
		//consoleLogger,
	)

	//TIP <h5>Create the CLI commands that can be executed</h5>
	createToDoCmd := create_todo_command.New(todoService, compositeLoggers)
	deleteToDoCmd := delete_todo_command.New(todoService, compositeLoggers)
	getToDoCmd := get_todo_command.New(todoService, compositeLoggers)
	listToDoCmd := list_todos_command.New(todoService, compositeLoggers)
	updatedToDoCmd := update_todo_command.New(todoService, compositeLoggers)
	makeToDoImportantCmd := make_todo_important_command.New(todoService, compositeLoggers)
	makeToDoNotImportantCmd := make_todo_not_important_command.New(todoService, compositeLoggers)
	cliCommands := []cli.ICommand{
		createToDoCmd,
		deleteToDoCmd,
		getToDoCmd,
		listToDoCmd,
		updatedToDoCmd,
		makeToDoImportantCmd,
		makeToDoNotImportantCmd,
	}
	rootCmd := cli.NewRootCommand(cliCommands)

	//TIP <h4>Execute the CLI</h4>
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
