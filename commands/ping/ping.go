package ping

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands/commandinterface"
	"go-dc-bot/utils/discord"
)

var pingCommand = commandinterface.CommandInterface{
	Exec:          exec,
	Name:          "ping",
	Alias:         []string{"pingbot"},
	CommandAccess: commandinterface.CommandAccessNone,
}

func exec(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
	discord.ChannelMessageSend(session, event.ChannelID, "pong")
}

func Get() *commandinterface.CommandInterface {
	return &pingCommand
}
