package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
}

func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening config file: %v", err)
	}
	defer file.Close()
	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		fmt.Printf("Error decoding config file: %v", err)
	}
	return &cfg
}
