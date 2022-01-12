package main

import (
	"fmt"
	"go-dc-bot/utils"
)

func main() {
	test2()
}

func test2() {
	flagTest := utils.GetConfigurationFlags()
	flagTestYaml := utils.GetConfigurationYaml()
	fmt.Println(flagTestYaml.BotToken)
	fmt.Println(*flagTest.ConfigPath)
}
