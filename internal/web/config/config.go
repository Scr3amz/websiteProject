package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"SERVER_PORT"`
	DBName string `mapstructure:"MYSQL_DB_NAME"`
	DBUser string `mapstructure:"MYSQL_USERNAME"`
	DBPass string `mapstructure:"MYSQL_PASSWORD"`
}

func LoadConfig() (Config, error){
	viper.AddConfigPath("./intrnal/web/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err!= nil {
        return Config{}, err
    }
	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}
