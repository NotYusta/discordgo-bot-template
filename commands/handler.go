package commands

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands/commandinterface"
	"go-dc-bot/commands/ping"
	"go-dc-bot/utils/discord"
)

var commandsData = make(commandinterface.CommandsData)

func Init() {
	println("[-------------------------------]")
	println("Registering commands..")
	temporaryCommandsData := []*commandinterface.CommandInterface{
		ping.Get(),
	}

	for _, commandElement := range temporaryCommandsData {
		println("Registering [" + commandElement.Name + "] command")
		commandsData[commandElement.Name] = commandElement

		for _, commandAlias := range commandElement.Alias {
			commandsData[commandAlias] = commandElement
		}
		println("Registered [" + commandElement.Name + "] command")
	}

	println("Succesfully registered all commands")
	println("[-------------------------------]")
}

func Exec(session *discordgo.Session, event *discordgo.MessageCreate, commandName string, args []string) {
	commandData := commandsData[commandName]
	if commandData == nil {
		discord.ChannelMessageSend(session, event.ChannelID, "Unknown command!")
		return
	}

	if commandData.CommandAccess == commandinterface.CommandAccessDJ {
		// DO DJ CHECKS
	}

	if commandData.CommandAccess == commandinterface.CommandAccessModerator {
		// DO MODERATOR CHECKS
	}

	if commandData.CommandAccess == commandinterface.CommandAccessAdmin {
		// DO ADMIN CHECKS
	}

	commandData.Exec(session, event, args)
}
