package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config store all configuration variables of the application
// all values are read by viper from a config file or env
type Config struct {
	DBDriver             string        `mapstructure:"GOOSE_DRIVER"`
	DBSource             string        `mapstructure:"GOOSE_DBSTRING"`
	HTTPServerURL        string        `mapstructure:"HTTP_SERVER_URL"`
	GRPCServerURL        string        `mapstructure:"GRPC_SERVER_URL"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
