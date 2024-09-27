package list_todos_command

import (
	"go.uber.org/fx"
	"todo_app/cli"
)

var Module = fx.Module("list_todos_command", fx.Provide(fx.Annotate(New, fx.As(new(cli.ICommand)), fx.ResultTags(`group:"commands"`))))
