package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	Port     uint16
	TimeZone string
}

type Db struct {
	Host    string
	Port    uint16
	User    string
	Passwd  string
	DBname  string
	SSLMode string
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal err reading config file: %v", err))
	}

	return Config{
		App: App{
			Port:     viper.GetUint16("app.server.port"),
			TimeZone: viper.GetString("app.tz"),
		},
		Db: Db{
			Host:    viper.GetString("database.host"),
			Port:    viper.GetUint16("database.port"),
			User:    viper.GetString("database.user"),
			Passwd:  viper.GetString("database.password"),
			DBname:  viper.GetString("database.dbname"),
			SSLMode: viper.GetString("database.sslmode"),
		},
	}
}
