package main

import (
	"zaixi.dev/leinao-docs-sync-bot/configs"
)

func main(){
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	println(config.Github.GithubWebhookSecret)
}