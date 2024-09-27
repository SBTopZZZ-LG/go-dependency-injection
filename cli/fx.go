package cli

import "go.uber.org/fx"

var Module = fx.Module("cli", fx.Provide(fx.Annotate(NewRootCommand, fx.ParamTags(`group:"commands"`))))
