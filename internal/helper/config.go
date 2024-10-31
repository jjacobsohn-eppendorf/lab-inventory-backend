package helper

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	var _, b, _, _ = runtime.Caller(0)
	viper.AddConfigPath(filepath.Join(filepath.Dir(b), "../../"))
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetDefault("PORT", "8081")
	err = viper.ReadInConfig()
	if err != nil {
		viper.AutomaticEnv()
	}

	err = viper.Unmarshal(&config)
	return
}
