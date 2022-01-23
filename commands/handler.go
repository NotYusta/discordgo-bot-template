package commands

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands/commandinterface"
	"go-dc-bot/commands/ping"
	"go-dc-bot/utils/discord"
	"sort"
)

var commandsData = make(commandinterface.CommandsData)

func Init() {
	println("[-------------------------------]")
	println("Registering commands..")
	temporaryCommandsData := []commandinterface.CommandInterface{
		ping.Get(),
	}

	for _, commandElement := range temporaryCommandsData {
		println("Registering [" + commandElement.Name + "] command")
		commandsData[commandElement.Name] = &commandElement

		for _, commandAlias := range commandElement.Alias {
			commandsData[commandAlias] = &commandElement
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

	if commandAccessContains(commandData.CommandAccess, commandinterface.CommandAccessAdmin) {
		// TODO: DO DJ CHECKS
		return
	}

	if commandAccessContains(commandData.CommandAccess, commandinterface.CommandAccessAdmin) {
		// TODO: DO MODERATOR CHECKS
		return
	}

	if commandAccessContains(commandData.CommandAccess, commandinterface.CommandAccessAdmin) {
		// TODO: DO ADMIN CHECKS
		return
	}

	commandData.Exec(session, event, args)
}

func commandAccessContains(a []commandinterface.CommandAccess, x commandinterface.CommandAccess) bool {
	accessLength := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	return accessLength > 0
}
