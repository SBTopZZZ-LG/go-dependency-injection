package update_todo_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
	"todo_app/cli"
	"todo_app/logger"
	"todo_app/services/to_do_service"
)

type UpdateTODOCommand struct {
	todoService to_do_service.ITODOService
	logger      logger.ILogger

	cli.ICommand
}

func New(todoService to_do_service.ITODOService, logger logger.ILogger) *UpdateTODOCommand {
	return &UpdateTODOCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *UpdateTODOCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "update_todo",
		Short: "Updates an existing todo using id",
		Args:  cobra.ExactArgs(2),
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *UpdateTODOCommand) run(_ *cobra.Command, args []string) {
	todoId, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		c.logger.Error("cannot parse to-do identifier", zap.Any("invalid_identifier", args[0]), zap.Error(err))
		fmt.Println("An error was encountered while trying to parse the to-do identifier")
		panic(err)
	}

	c.logger.Info("update to-do command", zap.Uint64("id", todoId))

	updatedTodoMessage := args[1]

	todoToUpdate, err := c.todoService.Get(todoId)
	if err != nil {
		c.logger.Error("cannot get to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to get the to-do\n%v\n", err)
		return
	}

	todoToUpdate.Message = updatedTodoMessage
	err = c.todoService.Update(todoToUpdate)
	if err != nil {
		c.logger.Error("cannot update to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to put to-do\n%v\n", err)
		return
	}

	fmt.Println("To-do was successfully updated.")
}
