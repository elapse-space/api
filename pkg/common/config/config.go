package config

import (
	"api/pkg/common/utils"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
}

func LoadConfig() (c Config, err error) {
	defer func() {
		if err == nil {
			utils.Logger.Info("Config loaded successfully")
		}
	}()
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
