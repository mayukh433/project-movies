package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUSERNAME  string `mapstructure:"DB_USERNAME"`
	DBPASSWORD  string `mapstructure:"DB_PASSWORD"`
	DBHOST      string `mapstructure:"DB_HOST"`
	DBPORT      string `mapstructure:"DB_PORT"`
	DBNAME      string `mapstructure:"DB_NAME"`
	HOST        string `mapstructure:"HOST"`
	SERVERPORT  string `mapstructure:"SERVER_PORT"`
	GATEWAYPORT string `mapstructure:"GATEWAY_PORT"`
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

	if err != nil {
		return
	}
	return config, nil
}
