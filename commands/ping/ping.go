package ping

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands/commandinterface"
)

var pingCommand = commandinterface.CommandInterface{
	Exec:          exec,
	Name:          "ping",
	Alias:         []string{"pingbot"},
	CommandAccess: commandinterface.CommandAccessNone,
}

func exec(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
	session.ChannelMessageSend(event.ChannelID, "pong")
}

func Get() (command *commandinterface.CommandInterface) {
	command = &pingCommand
	return
}
