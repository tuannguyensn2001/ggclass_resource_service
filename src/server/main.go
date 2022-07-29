package main

import (
	"ggclass_resource_service/src/cmd"
	"log"
)

func main() {
	rootCmd := cmd.GetRoot()

	err := rootCmd.Execute()

	if err != nil {
		log.Fatalln(err)
	}
}
