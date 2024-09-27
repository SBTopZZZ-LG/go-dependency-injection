# Go Todo DI 🚀

This project is a simple To-do application built with Go. While it covers basic features like marking tasks as important, the main goal here is to showcase how to implement **Dependency Inversion** and **Dependency Injection** in Go projects. We'll explore both manual methods and popular DI libraries like [**Uber fx**](https://github.com/uber-go/fx) and [**Google wire**](https://github.com/google/wire).

## Table of Contents

- [Branch Overview](#branch-overview)
- [About the `using_fx` Branch](#about-the-using_fx-branch)
  - [Perk 1: Eliminate Redundant Dependency Creations](#perk-1-eliminate-redundant-dependency-creations)
  - [Perk 2: Easier Dependency Management](#perk-2-easier-dependency-management)
  - [Perk 3: Dependency Lifecycle Management](#perk-3-dependency-lifecycle-management)
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

## About the `using_fx` Branch

The `using_fx` branch showcases our Go project enhanced with **Uber fx** to streamline **Dependency Injection** throughout the application's lifecycle. By leveraging **Uber fx**, we simplify dependency management, reduce boilerplate code, and enhance the modularity and maintainability of the codebase.

### Perk 1: Eliminate Redundant Dependency Creations

**Uber fx** effectively addresses the redundancy involved in creating dependencies across multiple locations in the codebase. With **fx**, you can define a module using `fx.Module` and provide constructors that return instances of dependencies, optionally handling errors. Check out [`utils/zap_util/fx.go`](./utils/zap_util/fx.go) to see how we define reusable **fx** modules.

**Example:** Consider a project with multiple CLI commands, where some commands depend on the Database. Without **fx**, loading the database configuration, creating the SQL connection, initiating the connection, and injecting the instance into an ORM would need to be replicated across all relevant commands. This leads to tight coupling between business logic and dependency constructors.

With **Uber fx**, you define these dependencies once in a module, eliminating the need to recreate them for each command and reducing coupling.

### Perk 2: Easier Dependency Management

Managing dependencies becomes significantly simpler with **fx**. Injecting dependencies into the application is streamlined, and thanks to the **Dependency Inversion** Principle, swapping between multiple implementations is straightforward. Additionally, **fx** allows binding concrete implementations to interfaces effortlessly. Explore [`services/to_do_service/fx.go`](./services/to_do_service/fx.go) or [`repositories/to_do_repository/fx.go`](./repositories/to_do_repository/fx.go) to see how we bind concrete types to abstract interfaces.

**Example Code Snippet from [`main.go`](./main.go):**

```go
fx.New(
    // . . . other modules

    // . . . Zap Logger Params
    zap_logger.Module, // Zap Logger Implementation
	
    // . . . Console Logger Params
    console_logger.Module, // Console Logger Implementation
	
    loggers.Module, // Composite Logger
    
    // . . . other modules
).Run()
```

In this snippet, loggers are injected into the application seamlessly. The `loggers` module acts as a composite logger, channeling logs to both `zap_logger.Module` and `console_logger.Module`. Managing loggers is effortless—simply remove or add logger modules in `main.go`, and **fx** automatically handles the injection of the logger slice into the composite logger.

### Perk 3: Dependency Lifecycle Management

**Uber fx**'s lifecycle hooks are invaluable for managing the lifecycle of dependencies, especially for lazy-initialized dependencies like database connections. Instead of manually initiating and terminating connections, **fx** allows you to define hooks that handle these processes automatically.

**Example with Database Connection:** When a database connection is created, it isn’t initiated immediately. Typically, you would manually initiate the connection and defer its termination. With **fx**, you can define lifecycle hooks that automatically initiate the connection when needed and gracefully terminate it when the application stops.
 
Refer to [`utils/sql_util/fx.go`](./utils/sql_util/fx.go) to see how we implement lifecycle hooks for the database connection, simplifying both initialization and cleanup processes.

### What's Next?

While manual **Dependency Injection** in the `with_dep_inv` branch offers improved structure and decoupling, managing numerous dependencies by hand can become cumbersome as the project grows. To maintain a simple, readable, and scalable codebase, leveraging **Dependency Injection** libraries like **Uber fx** and **Google wire** is highly beneficial. These libraries automate dependency wiring, reduce boilerplate code, and efficiently handle complex dependency graphs.

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
   git clone https://github.com/infraspec/go-todo-di.git
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

5. **Create Application Configuration**

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

6. **Create Logs Directory**

   To prevent the application from crashing due to missing log directories, create the `logs` folder:

   ```bash
   mkdir logs
   ```

7. **Run the Application**

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
