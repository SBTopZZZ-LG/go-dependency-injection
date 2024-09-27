package to_do_repository

import "go.uber.org/fx"

var Module = fx.Module("to_do_repository", fx.Provide(fx.Annotate(New, fx.As(new(ITODORepository)))))
