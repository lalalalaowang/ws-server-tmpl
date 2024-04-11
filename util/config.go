package util

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	LoginServer LoginServer
}

type LoginServer struct {
	Addr string
	Port string
}

var Conf *Config

func init() {
	err := cmdExecute()
	if err != nil {
		log.Fatal(err)
	}

	projectPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	configPath := projectPath + CC

	_, err = os.Stat(configPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(configPath, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
