package get_todo_command

import (
	"go.uber.org/fx"
	"todo_app/cli"
)

var Module = fx.Module("get_todo_command", fx.Provide(fx.Annotate(New, fx.As(new(cli.ICommand)), fx.ResultTags(`group:"commands"`))))
