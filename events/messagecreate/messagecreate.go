package messagecreate

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/eventinterface"
	"go-dc-bot/utils/config"
	"strings"
)

func Get() (eventHandler eventinterface.EventHandler) {
	eventHandler = eventinterface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.MessageCreate); ok {
			messageContent := event.Message.Content
			commandArgs := strings.Split(messageContent, " ")
			if strings.HasPrefix(commandArgs[0], config.GetConfigurationYaml().Options.Prefix) {
				commandArgs[0] = strings.Replace(commandArgs[0], "!!", "", 1)

				session.ChannelMessageSend(event.ChannelID, "Hello there! "+strings.Join(commandArgs, " "))

				// TODO: add command handlers here.
			}
		}
	}}

	return
}
