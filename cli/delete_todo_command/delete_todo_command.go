package delete_todo_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
	"todo_app/cli"
	"todo_app/logger"
	"todo_app/services/to_do_service"
)

type DeleteTODOCommand struct {
	todoService to_do_service.ITODOService
	logger      logger.ILogger

	cli.ICommand
}

func New(todoService to_do_service.ITODOService, logger logger.ILogger) *DeleteTODOCommand {
	return &DeleteTODOCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *DeleteTODOCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "delete_todo",
		Short: "Deletes a to-do using id",
		Args:  cobra.ExactArgs(1),
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *DeleteTODOCommand) run(_ *cobra.Command, args []string) {
	todoId, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		c.logger.Error("cannot parse to-do identifier", zap.Any("invalid_identifier", args[0]), zap.Error(err))
		fmt.Printf("An error was encountered while trying to parse the to-do identifier\n%v\n", err)
		return
	}

	c.logger.Info("delete to-do command", zap.Uint64("id", todoId))

	err = c.todoService.Delete(todoId)
	if err != nil {
		c.logger.Error("cannot delete to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to delete to-do\n%v\n", err)
		return
	}

	fmt.Println("To-do was successfully deleted.")
}
