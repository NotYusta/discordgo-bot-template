package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands"
	"go-dc-bot/events"
	"go-dc-bot/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	conf := utils.GetConfig().GetConfigurationYaml()
	botToken := conf.BotToken
	session, err := discordgo.New("Bot " + botToken)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
	if err != nil {
		println("An error occurred while starting the bot!")
		println(fmt.Sprint(err))
		return
	}

	commands.Init()
	events.Init(session)

	err = session.Open()
	if err != nil {
		println("An error occurred while starting the bot!")
		println(fmt.Sprint(err))
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signalResponse := <-sc
	println("Closing discord connection, received signal: " + signalResponse.String())

	err = session.Close()

}
