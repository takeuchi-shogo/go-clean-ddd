package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/constants"
)

type (
	DBConfig struct {
		Database    Database
		DBMS        string        `mapstructure:"dbms"`
		MaxIdleConn int           `mapstructure:"max_idle_conn"`
		MaxOpenConn int           `mapstructure:"max_open_conn"`
		MaxLifetime time.Duration `mapstructure:"max_life_time"`
	}
	Database struct {
		Host     string `mapstructure:"host"`
		UserName string `mapstructure:"user_name"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"database_name"`
		Port     string `mapstructure:"port"`
	}
)

func New(e string) *DBConfig {
	return newConfig(e)
}

func newConfig(e string) *DBConfig {
	config := &DBConfig{}
	viper.SetConfigType("yml")

	// viper.BindEnv("dbms", "db_config.dbms")
	// viper.BindEnv("max_idle_conn", "db_config.max_idle_conn")
	// viper.BindEnv("max_open_conn", "db_config.max_open_conn")
	// viper.BindEnv("max_life_time", "db_config.max_life_time")

	var fileName string
	switch e {
	case constants.Production:
		fileName = ""
	case constants.Local:
		fileName = "local"
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName(fileName)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("環境設定ファイルの読み込みに失敗しました")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return config
}
