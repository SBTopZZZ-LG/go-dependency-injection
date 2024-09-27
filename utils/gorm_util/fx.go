package gorm_util

import (
	"go.uber.org/fx"
)

var Module = fx.Module("gorm_util", fx.Provide(NewSilentGormInstanceWithMySQLDriver))
