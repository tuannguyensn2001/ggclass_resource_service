package main

import (
	"ggclass_resource_service/src/cmd"
	"ggclass_resource_service/src/config"
	"ggclass_resource_service/src/logger"
	"log"
)

func main() {

	config.Load()

	logger.InitLog()
	defer logger.SyncLog()

	rootCmd := cmd.GetRoot()

	err := rootCmd.Execute()

	if err != nil {
		log.Fatalln(err)
	}
}
