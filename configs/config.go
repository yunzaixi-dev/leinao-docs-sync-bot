package configs

import (
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Github struct {
		GithubWebhookSecret string `toml:"github_webhook_secret"`
	} `toml:"github"`
}

//加载配置文件
func LoadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile("configs/config.toml")
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	return config, err
}