package make_todo_important_command

import (
	"go.uber.org/fx"
	"todo_app/cli"
)

var Module = fx.Module("make_todo_important_command", fx.Provide(fx.Annotate(New, fx.As(new(cli.ICommand)), fx.ResultTags(`group:"commands"`))))
