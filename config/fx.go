package config

import "go.uber.org/fx"

var Module = fx.Module(
	"config",
	fx.Provide(
		func() (*Config, *DBConfig, *LoggerConfig, error) {
			conf, err := Load(DefaultConfigFileName)

			return conf, conf.DBConfig, conf.LoggerConfig, err
		},
	),
)
