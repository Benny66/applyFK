package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Url        string `yaml:"url"`
	Token      string `yaml:"token"`
	Ticket     string `yaml:"ticket"`
	ApplyPhone string `yaml:"applyPhone"`
}

var Conf *Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file failed, err: %v", err)
	}
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("unmarshal config file failed, err: %v", err)
	}
	if len(config.Url) == 0 {
		return fmt.Errorf("url is empty")
	}
	Conf = config
	fmt.Println("config loaded")
	fmt.Println(Conf)
	return nil
}
