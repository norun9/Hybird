package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	//Prd AppConfig `mapstructure:"prd"`
	Dev AppConfig `mapstructure:"dev"`
}

type AppConfig struct {
	DBConfig DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Host string `mapstructure:"host" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
	User string `mapstructure:"user" validate:"required"`
	Pass string `mapstructure:"pass" validate:"required"`
	Name string `mapstructure:"name" validate:"required"`
}

const AppName = "HYBIRD"

func Prepare() AppConfig {
	viper.SetEnvPrefix(AppName)

	viper.SetConfigName("config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	_, b, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(b)
	pkgDir := filepath.Dir(configDir)
	backendDir := filepath.Dir(pkgDir)

	viper.AddConfigPath(backendDir)
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	var appConfig AppConfig
	env := viper.GetString("env.name")
	switch env {
	case "dev":
		appConfig = c.Dev
		// TODO: Prd
	}
	return appConfig
}
