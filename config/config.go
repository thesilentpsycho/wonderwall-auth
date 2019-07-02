package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var config *GeneralConfig

func GetConfig() *GeneralConfig {
	if config != nil {
		return config
	}

	config = LoadConfig("config.yml")
	return config
}

type GeneralConfig struct {
	Environment           string
	ServiceName           string
	DatabaseHost          string
	DatabasePort          uint32
	DatabaseName          string
	LogLevel              string
	LogFilePath           string
	DefaultTimeout        time.Duration
	DefaultServerPort     uint32
	DefaultConfigFilePath string
}

func LoadConfig(filepath string) (a *GeneralConfig) {
	c := new(GeneralConfig)

	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return nil
	}

	c.Environment = viper.GetString("env")
	c.ServiceName = viper.GetString("service_name")
	c.DatabaseName = viper.GetString("database_name")
	c.DatabaseHost = viper.GetString("database_host")
	c.DatabasePort = viper.GetUint32("database_port")
	c.LogFilePath = viper.GetString("log_file")
	c.LogLevel = viper.GetString("log_level")
	c.DefaultTimeout = viper.GetDuration("default_timeout")
	c.DefaultServerPort = viper.GetUint32("default_server_port")
	c.DefaultConfigFilePath = viper.GetString("default_config_filepath")

	return c
}
