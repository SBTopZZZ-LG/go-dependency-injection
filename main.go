package main

import (
	"todo_app/config"
	"todo_app/entities"
)

func main() {
	//TIP <h5>Load Configuration from the `config.yaml` file</h5>
	conf, err := config.Load(config.DefaultConfigFileName)
	if err != nil {
		panic(err)
	}

	//TIP <h5>Create a GORM instance with the MySQL database connection</h5>
	gormDB, closeConn, err := InitGormWithMySQLDriver(conf.DBConfig)
	if err != nil {
		panic(err)
	}
	//TIP Make sure to close the database connection when `main` function exits
	defer closeConn()

	//TIP <h5>Run GORM Migration</h5>
	// This will create/update the To-do table automatically
	err = gormDB.AutoMigrate(&entities.TODO{})
	if err != nil {
		panic(err)
	}

	//TIP <h5>Create the To-do service</h5>
	todoService := InitTodoService(gormDB)

	//TIP <h5>Create the Logger</h5>
	compositeLoggers, err := InitCompositeLoggerWithZapAndConsole(conf.LoggerConfig)
	if err != nil {
		panic(err)
	}

	//TIP <h5>Create the CLI commands that can be executed</h5>
	rootCmd := InitCLICommands(todoService, compositeLoggers)

	//TIP <h4>Execute the CLI</h4>
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
