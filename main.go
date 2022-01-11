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

	fmt.Println(*flagTest.ConfigPath)
}
