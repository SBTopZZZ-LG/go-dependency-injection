package config

import (
	"github.com/spf13/viper"
)

type LoggerEncoderConfig struct {
	LineEnding string `mapstructure:"line_ending"`
}

type LoggerConfig struct {
	Level            string               `mapstructure:"level"`
	Development      bool                 `mapstructure:"development"`
	Encoding         string               `mapstructure:"encoding"`
	OutputPaths      []string             `mapstructure:"output_paths"`
	ErrorOutputPaths []string             `mapstructure:"error_output_paths"`
	EncoderConfig    *LoggerEncoderConfig `mapstructure:"encoder_config"`
}

type DBConfig struct {
	Driver   string `mapstructure:"driver"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Params   string `mapstructure:"params"`
}

var DefaultConfigFileName = "config.yaml"

type Config struct {
	LoggerConfig *LoggerConfig `mapstructure:"logger"`
	DBConfig     *DBConfig     `mapstructure:"database"`
}

func Load(filename string) (*Config, error) {
	viper.SetConfigFile(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
