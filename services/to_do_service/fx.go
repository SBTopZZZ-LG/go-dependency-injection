package to_do_service

import "go.uber.org/fx"

var Module = fx.Module("to_do_service", fx.Provide(fx.Annotate(New, fx.As(new(ITODOService)))))
