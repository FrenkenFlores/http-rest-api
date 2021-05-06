package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/FrenkenFlores/http-rest-api/internel/app/apiserver"
)

var config_path string

func init() {
	flag.StringVar(&config_path, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(config_path, config); err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
