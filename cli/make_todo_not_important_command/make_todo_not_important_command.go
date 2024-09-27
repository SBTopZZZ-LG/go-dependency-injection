package make_todo_not_important_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
	"todo_app/cli"
	"todo_app/logger"
	"todo_app/services/to_do_service"
)

type MakeTODONotImportantCommand struct {
	todoService to_do_service.ITODOService
	logger      logger.ILogger

	cli.ICommand
}

func New(todoService to_do_service.ITODOService, logger logger.ILogger) *MakeTODONotImportantCommand {
	return &MakeTODONotImportantCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *MakeTODONotImportantCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "make_todo_not_important",
		Short: "Marks a to-do as not important using id",
		Args:  cobra.ExactArgs(1),
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *MakeTODONotImportantCommand) run(_ *cobra.Command, args []string) {
	todoId, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		c.logger.Error("cannot parse to-do identifier", zap.Any("invalid_identifier", args[0]), zap.Error(err))
		fmt.Printf("An error was encountered while trying to parse the to-do identifier\n%v\n", err)
		return
	}

	c.logger.Info("make to-do not important command", zap.Uint64("id", todoId))

	todoToUpdate, err := c.todoService.Get(todoId)
	if err != nil {
		c.logger.Error("cannot get to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to get the to-do\n%v\n", err)
		return
	}

	todoToUpdate.IsImportant = false
	err = c.todoService.Update(todoToUpdate)
	if err != nil {
		c.logger.Error("cannot update to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to mark the to-do as important\n%v\n", err)
		return
	}

	fmt.Printf("\n%v\n\n", todoToUpdate)
}
