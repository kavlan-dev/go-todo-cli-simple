package config

import (
	"fmt"
	"os"
)

type Config struct {
	TaskFile string
}

func InitConfig() (*Config, error) {
	var config Config
	config.TaskFile = envOrDefault("TASK_FILE", "tasks.json")

	if config.TaskFile == "" {
		return nil, fmt.Errorf("TASK_FILE не указан")
	}

	return &config, nil
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}

	return value
}
