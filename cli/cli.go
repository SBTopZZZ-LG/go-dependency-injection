package cli

import (
	"github.com/spf13/cobra"
)

type ICommand interface {
	Init(rootCmd *cobra.Command)
}

type Command struct {
	commands []ICommand

	rootCmd *cobra.Command
}

func NewRootCommand(commands []ICommand) *Command {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "root - a simple To-do management application",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return &Command{
		commands: commands,
		rootCmd:  rootCmd,
	}
}

func (c *Command) Execute() error {
	for _, cmd := range c.commands {
		cmd.Init(c.rootCmd)
	}

	if err := c.rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
