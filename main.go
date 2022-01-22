package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events"
	"go-dc-bot/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	BotToken := utils.GetConfig().GetConfigurationYaml().BotToken
	discord, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		println("An error occurred while starting the bot!")
		println(fmt.Sprint(err))
		return
	}

	events.Init(discord)
	err = discord.Open()
	if err != nil {
		println("An error occurred while starting the bot!")
		println(fmt.Sprint(err))
		return
	}
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signalResponse := <-sc
	println("Stopping server, received signal: " + signalResponse.String())

	err = discord.Close()

}
