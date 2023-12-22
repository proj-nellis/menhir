package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database string  `yaml:"database"`
	NodeId   int64   `yaml:"node_id"`
	Mode     *string `yaml:"mode"`
}

func Init() (cfg Config, err error) {
	config := Config{}
	log.Print("Loading config file ...")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory.")
		return config, err
	}
	log.Println(pwd)
	config_bytes, err := ioutil.ReadFile(pwd + "/config.yml")
	if err != nil {
		log.Fatal("Failed to read config file.")
		return config, err
	}

	yaml.Unmarshal(config_bytes, &config)
	return config, nil
}
