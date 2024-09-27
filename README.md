# go-todo-app

## Project Architecture

The project is a To-Do application structured into several key components:

1. **Entities**: Contains the data models, such as `TODO`.
2. **Repositories**: Handles data persistence and retrieval operations.
3. **Services**: Contains the business logic of the application.
4. **CLI Commands**: Provides a command-line interface for interacting with the application.
5. **Main**: The entry point of the application.

### Business Logic

The business logic of the application is encapsulated within the `services` layer. This layer interacts with the `repositories` to perform CRUD operations on the To-Do items. The `services` layer ensures that the business rules and validations are applied before any data manipulation occurs.

## Dependency Injection with Uber Fx

Uber Fx is a dependency injection framework for [Go](https://go.dev/). It helps manage dependencies in a type-safe manner, reducing boilerplate code and improving testability.

### Benefits of Using Uber Fx

1. **Type Safety**: Ensures that dependencies are correctly provided at compile time.
2. **Reduced Boilerplate**: Automatically generates the wiring code, reducing manual setup.
3. **Improved Testability**: Makes it easier to inject mock dependencies for testing.
4. **Maintainability**: Changes in dependencies are easier to manage.

### Downsides of Using Uber Fx

1. **Learning Curve**: Requires understanding of how Fx works.
2. **Runtime Overhead**: Introduces some runtime overhead for dependency resolution.
3. **Complexity**: Can add complexity to the project setup.

### How Uber Fx is Used in the Project

To use Uber Fx, you typically define modules that include constructors for your dependencies. Here’s a simplified example of how it is set up in the project:

```go
package main

import (
	"go.uber.org/fx"
	"your_project_path/repositories/to_do_repository"
	"your_project_path/services/to_do_service"
)

func main() {
	fx.New(
		to_do_repository.Module,
		to_do_service.Module,
		fx.Invoke(func () {}),
	).Run()
}
```

#### Explanation:

- **fx.Provide:** Used to specify how to construct the dependencies.
- **fx.Invoke:** Used to specify functions that should be executed on application start.

### Detailed Explanation

1. **Provider Functions:** These are functions that return instances of the required types. For example, `NewTODORepository` and `NewTODOService` are provider functions that return instances of `TODORepository` and `TODOService`, respectively.
2. **Invoke Functions:** These are functions that Fx will call to initialize your application. In this case, initialize is the function that will be called to set up the `TODOService` with all its dependencies properly wired.
3. **Fx New:** The `fx.New` function is used to create a new Fx application, specifying the provider and invoke functions.

## Additional Notes

- **Type Safety:** One of the primary advantages of using Uber Fx is that it ensures type safety at runtime. This means that any issues with dependency injection will be caught during application startup rather than at runtime.
- **Reduced Boilerplate:** By automatically generating the wiring code, Uber Fx significantly reduces the amount of boilerplate code that developers need to write.
- **Improved Testability:** With Uber Fx, it becomes easier to inject mock dependencies for testing purposes, thereby improving the testability of the code.
- **Maintainability:** Changes in dependencies are easier to manage with Uber Fx. If a new dependency is added or an existing one is modified, the changes need to be made only in the provider functions, and Fx will take care of the rest.

In conclusion, the use of Uber Fx in this To-Do application enhances type safety, reduces boilerplate, and improves maintainability and testability. However, it does introduce some complexity and a learning curve, which should be considered when deciding to use it in a project.
