package config

import "github.com/spf13/viper"

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Config struct {
	DBConfig DBConfig
}

func LoadConfig() Config {
	var config Config

	viper.SetEnvPrefix("HYBIRD")
	viper.AutomaticEnv()

	config.DBConfig.Pass = viper.GetString("DB_PASS")
	config.DBConfig.Port = viper.GetString("DB_PORT")
	config.DBConfig.Host = viper.GetString("DB_HOST")
	config.DBConfig.Name = viper.GetString("DB_NAME")
	config.DBConfig.User = viper.GetString("DB_USER")
	return config
}
