package cli

import (
	"github.com/spf13/cobra"
)

type ICommand interface {
	Init(rootCmd *cobra.Command)
}

type RootCommand struct {
	commands []ICommand

	rootCmd *cobra.Command
}

func NewRootCommand(commands []ICommand) *RootCommand {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "root - a simple To-do management application",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return &RootCommand{
		commands: commands,
		rootCmd:  rootCmd,
	}
}

func (c *RootCommand) Execute() error {
	for _, cmd := range c.commands {
		cmd.Init(c.rootCmd)
	}

	if err := c.rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
