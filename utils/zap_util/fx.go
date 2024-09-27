package zap_util

import (
	"go.uber.org/fx"
)

var Module = fx.Module("zap_util", fx.Provide(NewZapLogger))
