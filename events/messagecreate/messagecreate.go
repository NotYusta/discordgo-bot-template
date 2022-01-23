package messagecreate

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/commands"
	"go-dc-bot/events/eventinterface"
	"go-dc-bot/utils"
	"go-dc-bot/utils/config"
	"strings"
)

func Get() eventinterface.EventHandler {
	return eventinterface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.MessageCreate); ok {
			messageContent := event.Message.Content
			commandArgs := strings.Split(messageContent, " ")
			if strings.HasPrefix(commandArgs[0], config.GetConfigurationYaml().Options.Prefix) {
				commandName := strings.Replace(commandArgs[0], "!!", "", 1)
				commandArgs, _ := utils.RemoveArrayFromIndex(commandArgs, 0)

				commands.Exec(session, event, commandName, commandArgs)
			}
		}
	}}
}
