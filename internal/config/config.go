package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	data, err := os.ReadFile("../gatorconfig.json")
	if err != nil {
		return Config{}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}
	}

	return config
}

func (c *Config) SetUser() {}
