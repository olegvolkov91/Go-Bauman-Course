package api

import "github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/storage"

type Config struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}
}
