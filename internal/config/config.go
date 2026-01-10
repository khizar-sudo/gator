package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFilePath = "./gatorconfig.json"

func Read() Config {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}
	}

	return config
}

func (c *Config) SetUser() error {
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder((file))
	return encoder.Encode(c)
}
