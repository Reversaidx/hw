package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	Http    struct {
		Host string
		Port int
	}

	// TODO
}
type LoggerConf struct {
	Level    string
	log_file string
}
type StorageConf struct {
	connectionString string
	Type             string
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
	fmt.Println(Conf.Http.Host)
	fmt.Println(viper.AllSettings())

	return Conf
}
