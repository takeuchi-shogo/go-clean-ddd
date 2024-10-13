package infrastructure

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/constants"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		UserName string `mapstructure:"user_name"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"database_name"`
	}
	Server struct {
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	}
}

func NewConfig(env string) *Config {
	return newConfig(env)
}

func newConfig(e string) *Config {
	config := &Config{}
	viper.SetConfigType("yml")

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
