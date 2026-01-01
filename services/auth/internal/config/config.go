package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env     string      `mapstructure:"env"`
	GRPC    GRPCConfig  `mapstructure:"grpc"`
	Redis   RedisConfig `mapstructure:"redis"`
	JWT     JWTConfig   `mapstructure:"jwt"`
	Clients struct {
		UsersService struct {
			Address string `mapstructure:"address"`
		} `mapstructure:"users_service"`
	} `mapstructure:"clients"`
}

type GRPCConfig struct {
	Port string `mapstructure:"port"`
}

func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
