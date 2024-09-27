# Go Todo DI 🚀

This project is a simple To-do application built with Go. While it covers basic features like marking tasks as important, the main goal here is to showcase how to implement **Dependency Inversion** and **Dependency Injection** in Go projects. We'll explore both manual methods and popular DI libraries like [**Uber fx**](https://github.com/uber-go/fx) and [**Google wire**](https://github.com/google/wire).

## Table of Contents

- [Branch Overview](#branch-overview)
- [About the `without_dep_inv` Branch](#about-the-without_dep_inv-branch)
  - [What's Problematic Here?](#whats-problematic-here)
  - [Why You Should Care](#why-you-should-care)
  - [Next Steps](#next-steps)
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

## About the `without_dep_inv` Branch

This branch represents our initial Go project setup with basic dependency management.

### What's Problematic Here?

1. **Tight Coupling Between Repository and Service**
   - **Why it’s bad:** Makes it hard to swap out implementations or reuse components. Testing becomes a headache because dependencies are concrete and not easily mockable.

2. **Tight Coupling Between CLI Commands, Service, and Zap Logger**
   - **Why it’s bad:** Besides design rigidity and testing difficulties, having direct ties to external libraries like Zap means any updates or changes can lead to costly migrations.

### Why You Should Care

Without **Dependency Inversion** and **Dependency Injection**, your codebase can quickly become rigid and difficult to maintain. It hampers flexibility, modularity, and makes automated testing a nightmare.

### Next Steps

Ready to see improvements? Check out the [`with_dep_inv`](https://github.com/infraspecdev/go-todo-di/tree/with_dep_inv) branch to see how we tackle these issues by introducing Dependency Inversion. From there, you can explore `using_fx` and `using_wire` to see how different DI libraries can streamline your project.

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
   git clone https://github.com/infraspecdev/go-todo-di.git
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
