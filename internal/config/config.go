package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
	}

	App struct {
		LogLevel        string  `env-required:"true" yaml:"log_level" env:"SS_LOG_LEVEL"`
		BotToken        string  `env-required:"true" yaml:"bot_token" env:"SS_BOT_TOKEN"`
		BotUsername     string  `env-required:"true" yaml:"bot_username" env:"SS_BOT_USERNAME"`
		BotAllowedUsers []int64 `env-required:"true" yaml:"bot_allowed_users" env:"SS_BOT_ALLOWED_USERS" env-delim:","`
		// Ignore   []string `env-required:"true" yaml:"ignore"  env:"BLOSSOM_TW_IGNORE" env-delim:","`
		WebHookURL    string `env-required:"true" yaml:"webhook_url" env:"SS_WEBHOOK_URL"`
		JackettURL    string `env-required:"true" yaml:"jackett_url" env:"SS_JACKETT_URL"`
		JackettApiKey string `env-required:"true" yaml:"jackett_api_key" env:"SS_JACKETT_API_KEY"`
		RedisAddr     string `env-required:"true" yaml:"redis_addr" env:"SS_REDIS_ADDR"`
	}
)

// New returns app config.
func New() (*Config, error) {
	c := &Config{}

	err := cleanenv.ReadEnv(c)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return c, nil
}
