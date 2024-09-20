package get_todo_command

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
	"todo_app/cli"
	"todo_app/services/to_do_service"
)

type GetTODOCommand struct {
	todoService *to_do_service.TODOService
	logger      *zap.Logger

	cli.ICommand
}

func New(todoService *to_do_service.TODOService, logger *zap.Logger) *GetTODOCommand {
	return &GetTODOCommand{
		todoService: todoService,
		logger:      logger,
	}
}

func (c *GetTODOCommand) Init(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "get_todo",
		Short: "Gets to-do using id",
		Args:  cobra.ExactArgs(1),
		Run:   c.run,
	}

	rootCmd.AddCommand(cmd)
}

func (c *GetTODOCommand) run(_ *cobra.Command, args []string) {
	todoId, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		c.logger.Error("cannot parse to-do identifier", zap.Any("invalid_identifier", args[0]), zap.Error(err))
		fmt.Printf("An error was encountered while trying to parse the to-do identifier\n%v\n", err)
		return
	}

	c.logger.Info("get to-do command", zap.Uint64("id", todoId))

	todo, err := c.todoService.Get(todoId)
	if err != nil {
		c.logger.Error("cannot get to-do", zap.Error(err))
		fmt.Printf("An error was encountered while trying to get the to-do\n%v\n", err)
		return
	}

	fmt.Printf("\n%v\n\n", todo)
}
