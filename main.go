package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot " + utils.GetConfigurationYaml().BotToken)

	if err != nil {
		println("An error occurred while starting the bot!")
		println(fmt.Sprint(err))
		return
	}

	discord.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		println(discord.Token)
	})
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
