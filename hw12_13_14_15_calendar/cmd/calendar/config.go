package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	HTTP    struct {
		Host string
		Port int
	}

	// TODO
}

type LoggerConf struct {
	Level string
}

type StorageConf struct {
	// connectionString string
	Type string
}

func NewConfig(configFile string) Config {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	fmt.Println(viper.ConfigFileUsed())
	viper.ReadInConfig()
	fmt.Println(viper.Get("logger"))
	Conf := Config{}
	viper.Unmarshal(&Conf)
	fmt.Println(Conf.HTTP.Host)
	fmt.Println(viper.AllSettings())

	return Conf
}
