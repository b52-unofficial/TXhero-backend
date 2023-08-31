package config

import (
	"log"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	App      App
	DataBase Database
	Secret   string
	LogLevel string
	Eth      Eth
	Job      Job
}

type App struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Eth struct {
	Rpc string
}

type Job struct {
	SyncTx              string
	NextRoundWinningBid string
}

func newConfig() {
	config = new(Config)

	v := NewViper()

	if err := v.Unmarshal(config); err != nil {
		log.Printf("cannot parse config file\n")
	}
}

func GetConfig() *Config {
	if config == nil {
		once.Do(newConfig)
	}
	return config
}
