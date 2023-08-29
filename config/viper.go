package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

const (
	ConfigFileName = ".config"
	ConfigType     = "yaml"
)

func init() {
	Viper = NewViper()
}

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(ConfigFileName)
	v.SetConfigType(ConfigType)
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return v
}
