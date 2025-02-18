package configs

import (
	"os"
	"encoding/json"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Github struct {
		GithubWebhookSecret string `toml:"github_webhook_secret"`
		GithubRepoFrom      string `toml:"github_repo_from"`
		GithubRepoTo        string `toml:"github_repo_to"`
	} `toml:"github"`
}

func (c Config) String() string {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(data)
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