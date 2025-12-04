# Go Todo DI 🚀

This project is a simple To-do application built with Go. While it covers basic features like marking tasks as important, the main goal here is to showcase how to implement **Dependency Inversion** and **Dependency Injection** in Go projects. We'll explore both manual methods and popular DI libraries like [**Uber fx**](https://github.com/uber-go/fx) and [**Google wire**](https://github.com/google/wire).

## Table of Contents

- [Branch Overview](#branch-overview)
- [About the `using_wire` Branch](#about-the-using_wire-branch)
  - [Perk 1: Eliminate Redundant Dependency Creation](#perk-1-eliminate-redundant-dependency-creation)
  - [Perk 2: No Underlying Framework Overhead](#perk-2-no-underlying-framework-overhead)
  - [Perk 3: Compile-Time Safety](#perk-3-compile-time-safety)
  - [What's Next?](#whats-next)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setting Up the SQL Database](#setting-up-the-sql-database)
    - [Using Docker to Run MySQL](#using-docker-to-run-mysql)
  - [Running the Application](#running-the-application)
- [Usage](#usage)
  - [Create a To-do](#create-a-to-do)
  - [List To-dos](#list-to-dos)
  - [Get a To-do by ID](#get-a-to-do-by-id)
  - [Update a To-do](#update-a-to-do)
  - [Delete a To-do](#delete-a-to-do)
  - [Mark a To-do as Important](#mark-a-to-do-as-important)
  - [Mark a To-do as Not Important](#mark-a-to-do-as-not-important)

## Branch Overview

This repository has four branches, each highlighting different approaches to dependency management:

- **`without_dep_inv`** (default): The starting point with basic dependency injection, without dependency inversion.
- **`with_dep_inv`**: Introduces Dependency Inversion to address the initial shortcomings.
- **`using_fx`**: Implements Dependency Injection using the Uber fx library.
- **`using_wire`**: Utilizes Google wire for Dependency Injection.

## About the `using_wire` Branch

The `using_wire` branch demonstrates our Go project utilizing **Google Wire** for **Dependency Injection**. **Google Wire** is a compile-time tool that automates the wiring of dependencies, resulting in cleaner code without adding a runtime framework.

### Perk 1: Eliminate Redundant Dependency Creation

Similar to **Uber fx**, **Google Wire** helps eliminate redundancy in dependency creation. By logically grouping dependencies, we can simplify the instantiation of high-level components.

**Example:** By grouping the SQL utility dependency with the GORM utility, we can directly instantiate the GORM instance. The SQL connection is automatically created and injected into GORM without manual intervention.

**Before:**

```go
func main() {
	// . . .
	
	sqlDB, err := sql_util.CreateConnection(conf.DBConfig)
	// . . .
	
	gormDB, err := gorm_util.NewSilentGormInstanceWithMySQLDriver(sqlDB)
	// . . .
	
	todoRepository := to_do_repository.New(gormDB)
	todoService := to_do_service.New(todoRepository)
	
	zap, err := zap_util.NewZapLogger(conf.LoggerConfig)
	// . . .
	zapLogger := zap_logger.New(zap)
	consoleLogger := console_logger.New(2)
	compositeLoggers := loggers.New(
		zapLogger,
		consoleLogger,
	)
	
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
	
	// . . .
}
```

**After:**

```go
func main() {
	// . . .
	
	gormDB, closeConn, err := InitGormWithMySQLDriver(conf.DBConfig)
	// . . .
	
	todoService := InitTodoService(gormDB)
	
	compositeLoggers, err := InitCompositeLoggerWithZapAndConsole(conf.LoggerConfig)
	// . . .
	
	rootCmd := InitCLICommands(todoService, compositeLoggers)
	
	// . . .
}
```

### Perk 2: No Underlying Framework Overhead

One significant advantage of using **Google Wire** is that it doesn't add an underlying framework to your application. Since **Wire** generates plain Go code during compile time, there's:

1. **Minimal Learning Curve:** No extensive framework APIs to learn. If you're familiar with Go, you can understand and work with the generated code easily.
2. **Performance Benefits:** Without the overhead of a runtime framework, your application may perform better.
3. **Ease of Debugging:** The generated code is straightforward Go code, making it easier to debug and trace execution flow.

Because Wire generates explicit code for dependency injection, it makes the wiring of your application transparent. This explicitness can make the codebase more maintainable, especially for developers who prefer less magic and more clarity in how dependencies are resolved.

**Example:**

`wire.go`:

```go
//go:build wireinject
// +build wireinject

package main

import (
    // . . .
)

func CreateC(id int, label string) (*C, error) {
	wire.Build(
		NewA,
		NewB,
		NewC,
	)
	
	return nil, nil
}
```

`wire_gen.go`:

```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
    // . . .
)

func CreateC(id int, label string) (*C, error) {
	a := NewA(id)
	b, err := NewB(a, message)
	if err != nil {
		return nil, err
	}
	c, err := NewC(b)
	if err != nil {
		return nil, err
	}
	return c, nil
}
```

### Perk 3: Compile-Time Safety

Google Wire performs dependency injection at compile time, which allows for early detection of missing dependencies or configuration errors. This means:

1. **Early Error Detection:** Dependency issues are caught during compilation, reducing runtime surprises.
2. **Type Safety:** Ensures that the correct types are injected, preventing type mismatch errors.

### What's Next?

While **Google Wire** provides a compile-time, framework-free approach to **Dependency Injection**, it's essential to choose the right tool based on your project's needs:

- **Explore the [`using_fx` branch](https://github.com/SBTopZZZ-LG/go-dependency-injection/tree/using_fx)** to see how **Uber fx** offers a runtime framework with additional features like lifecycle management and built-in components.
- **Consider Your Project's Complexity:** For projects where runtime flexibility and advanced features are required, a framework like **Uber fx** might be more suitable.
- **Evaluate Performance Needs:** If minimal runtime overhead and compile-time safety are priorities, **Google Wire** is an excellent choice.

## Getting Started

### Prerequisites

- **Go**: Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
- **Docker**: For running the MySQL database. You can download it from [docs.docker.com/engine/install](https://docs.docker.com/engine/install/).

### Setting Up the SQL Database

To run the Go Todo DI application, you'll need a MySQL database. Using Docker simplifies the setup process.

#### Using Docker to Run MySQL

1. **Pull the MySQL Docker Image**
   
   ```bash
   docker pull mysql:latest
   ```

2. **Start the MySQL Container**

   Replace `your_password` with a secure password of your choice.

   ```bash
   docker run --name go-todo-mysql -e MYSQL_ROOT_PASSWORD=your_password -e MYSQL_DATABASE=todo_db -p 3306:3306 -d mysql:latest
   ```

3. **Verify the MySQL Container is Running**

   ```bash
   docker ps | grep go-todo-mysql
   ```
   
   You should see an entry for `go-todo-mysql` in the list of running containers.

### Running the Application

1. **Clone the Repository**

   ```bash
   git clone https://github.com/SBTopZZZ-LG/go-dependency-injection.git
   ```

2. **Navigate to the Project Directory**

   ```bash
   cd go-todo-di
   ```

3. **Checkout the `without_dep_inv` Branch**

   ```bash
   git checkout without_dep_inv
   ```

4. **Install dependencies**

   ```bash
   go mod tidy
   ```

5. **Generate Google Wire Boilerplate**

   ```bash
   go run github.com/google/wire/cmd/wire
   ```

6. **Create Application Configuration**

   Generate the `config.yaml` file with the necessary configurations:

   ```bash
   cat <<'EOF' > config.yaml
   logger:
     level: "info"
     development: true
     encoding: "json"
     output_paths:
       - "./logs/all-logs.log"
     error_output_paths:
       - "./logs/errors.log"
     encoder_config:
       line_ending: "\n"
   
   database:
     driver: "mysql"
     user: "root"
     password: "your_password"
     host: "localhost"
     port: 3306
     name: "todo_db"
     params: "parseTime=true"
   EOF
   ```
   
   **Note: Replace your_password with the password you set for the MySQL root user.**

7. **Create Logs Directory**

   To prevent the application from crashing due to missing log directories, create the `logs` folder:

   ```bash
   mkdir logs
   ```

8. **Run the Application**

   ```bash
   go run main.go
   ```
   
   The application should now be running, and you can interact with the CLI to manage your To-dos.

## Usage

Interact with the Go Todo DI application using the following commands. Each command allows you to manage your to-dos effectively through the CLI.

```bash
go run main.go <command> [{<id>|<content>|<id> <content>}]
```

### Create a To-do

**Syntax**

```bash
go run main.go create_todo <message>
```

**Example**

```bash
go run main.go create_todo "Write a Blog"
```

### List To-dos

**Syntax**

```bash
go run main.go list_todos
```

### Get a To-do by ID

Retrieve a specific to-do using its unique ID.

**Syntax**

```bash
go run main.go get_todo <id>
```

**Example**

```bash
go run main.go get_todo 1
```

### Update a To-do

**Syntax**

```bash
go run main.go update_todo <id> <new_message>
```

**Example**

```bash
go run main.go update_todo 1 "Buy groceries and cook dinner"
```

### Delete a To-do

**Syntax**

```bash
go run main.go delete_todo <id>
```

**Example**

```bash
go run main.go delete_todo 1
```

### Mark a To-do as Important

Flag a to-do as important to prioritize it.

**Syntax**

```bash
go run main.go make_todo_important <id>
```

**Example**

```bash
go run main.go make_todo_important 1
```

### Mark a To-do as Not Important

Remove the important flag from a to-do.

**Syntax**

```bash
go run main.go make_todo_not_important <id>
```

**Example**

```bash
go run main.go make_todo_not_important 1
```
