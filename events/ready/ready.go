package ready

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/eventinterface"
	"go-dc-bot/utils"
	"go-dc-bot/utils/config"
	"strings"
)

func Get() (eventHandler eventinterface.EventHandler) {
	eventHandler = eventinterface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.Ready); ok {
			var config = utils.GetConfig().GetConfigurationYaml()
			println("Logged in as " + event.User.Username)
			if config.Options.Activity.Enable {
				activityOptions := config.Options.Activity
				switch strings.ToUpper(activityOptions.Type) {
				case "PLAYING":
					setStatusData(session, config, 0)
					break
				case "STREAMING":
					setStatusData(session, config, 1)
					break
				case "LISTENING":
					setStatusData(session, config, 2)
					break
				case "WATCHING":
					setStatusData(session, config, 3)
					break
				case "CUSTOM":
					setStatusData(session, config, 4)
					break
				case "COMPETING":
					setStatusData(session, config, 5)
					break

				}
			}

		}
	}}

	return eventHandler
}

func setStatusData(session *discordgo.Session, config *config.ConfStructure, activityType discordgo.ActivityType) {
	activityOptions := config.Options.Activity
	err := session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{{
			Name: activityOptions.Description,
			Type: activityType,
			URL:  activityOptions.Url,
		}},
	})

	sendErrMessage(err, "Cannot set the status!")
}

func sendErrMessage(value error, message string) {
	if value != nil {
		println(message)
	}
}
