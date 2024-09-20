package list_todos_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"todo_app/cli"
	"todo_app/services/to_do_service"
)

type ListTODOsCommand struct {
	todoService *to_do_service.TODOService
	logger      *zap.Logger

	cli.ICommand
}

func New(todoService *to_do_service.TODOService, logger *zap.Logger) *ListTODOsCommand {
	return &ListTODOsCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *ListTODOsCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "list_todos",
		Short: "Lists all to-dos",
		Args:  cobra.NoArgs,
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *ListTODOsCommand) run(_ *cobra.Command, _ []string) {
	todos, err := c.todoService.List()
	if err != nil {
		c.logger.Error("cannot list the to-dos", zap.Error(err))
		fmt.Printf("An error was encountered while trying to list the to-dos\n%v\n", err)
		return
	}

	c.logger.Info("list to-dos command")

	for _, todo := range todos {
		fmt.Printf("\n%v\n", todo)
	}
	fmt.Println()
}
