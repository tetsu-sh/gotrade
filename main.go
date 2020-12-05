package main

import (
	"fmt"
	"gotrading/config"
	"gotrading/config/bitflyer"
	"gotrading/utils"
	"log"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
