package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	cfg := api.NewConfig()

	_, err := toml.DecodeFile(configPath, cfg) // Десериализация toml файла
	if err != nil {
		log.Println("can not find configs file. using default values", err)
	}

	server := api.New(cfg)
	log.Fatal(server.Start())
}
