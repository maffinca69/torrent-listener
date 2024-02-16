package infrastructure

import (
	"encoding/json"
	"os"
)

type RepositoryConfig struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type TelegramConfig struct {
	ChatID                   int64  `json:"chat_id"`
	DownloadCompletedMessage string `json:"download_completed_message"`
	WelcomeMessage           string `json:"welcome_message"`
}

type Configuration struct {
	Repository     []RepositoryConfig `json:"repositories"`
	TelegramConfig TelegramConfig     `json:"telegram"`
	CronExpression string             `json:"cron_expression"`
}

const ConfigName = "config.json"

var configInstance *Configuration

func Config() *Configuration {
	if configInstance == nil {
		configInstance = setupConfig()
	}

	return configInstance
}

func setupConfig() *Configuration {
	content, _ := os.ReadFile(ConfigName)
	payload := &Configuration{}

	if err := json.Unmarshal(content, &payload); err != nil {
		panic("Error load config " + ConfigName)
	}
	return payload
}
