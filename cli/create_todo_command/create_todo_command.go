package create_todo_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"todo_app/cli"
	"todo_app/entities"
	"todo_app/services/to_do_service"
)

type CreateTODOCommand struct {
	todoService *to_do_service.TODOService
	logger      *zap.Logger

	cli.ICommand
}

func New(todoService *to_do_service.TODOService, logger *zap.Logger) *CreateTODOCommand {
	return &CreateTODOCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *CreateTODOCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "create_todo",
		Short: "Creates a new to-do",
		Args:  cobra.ExactArgs(1),
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *CreateTODOCommand) run(_ *cobra.Command, args []string) {
	todoMessage := args[0]

	c.logger.Info("create todo command", zap.String("message", todoMessage))

	todo := entities.NewTODO(todoMessage, false)

	err := c.todoService.Create(todo)
	if err != nil {
		c.logger.Error("cannot create to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to create to-do\n%v\n", err)
		return
	}

	fmt.Println("To-do was successfully created.")
}
