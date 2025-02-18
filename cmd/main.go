package main

import (
	"zaixi.dev/leinao-docs-sync-bot/configs"
	"zaixi.dev/leinao-docs-sync-bot/internal/utils/logger"
)

func main(){
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	logger.InitLogger()
	logger.Info("config: " + config.String())
}