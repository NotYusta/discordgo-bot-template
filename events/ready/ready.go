package ready

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/eventinterface"
	"go-dc-bot/utils"
	"strings"
)

var (
	presenceOptions = utils.GetConfig().GetConfigurationYaml().Options.Presence
	activityOptions = utils.GetConfig().GetConfigurationYaml().Options.Activity
)

func Get() eventinterface.EventHandler {
	return eventinterface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.Ready); ok {

			println("Logged in as " + event.User.Username)

			if activityOptions.Enable {
				switch strings.ToUpper(activityOptions.Type) {
				case "PLAYING":
					setStatusDataWithPresence(session, 0)
					break
				case "STREAMING":
					setStatusDataWithPresence(session, 1)
					break
				case "LISTENING":
					setStatusDataWithPresence(session, 2)
					break
				case "WATCHING":
					setStatusDataWithPresence(session, 3)
					break
				case "CUSTOM":
					setStatusDataWithPresence(session, 4)
					break
				case "COMPETING":
					setStatusDataWithPresence(session, 5)
					break
				}
			} else if presenceOptions.Enable {
				err := session.UpdateStatusComplex(discordgo.UpdateStatusData{
					Status: strings.ToLower(presenceOptions.Type),
				})

				sendErrMessage(err, "Cannot set the activity status!")
			}
		}
	}}
}

func setStatusDataWithPresence(session *discordgo.Session, activityType discordgo.ActivityType) {
	if presenceOptions.Enable {
		setStatusData(session, presenceOptions.Type, activityType)
	} else {
		setStatusData(session, "online", activityType)
	}
}
func setStatusData(session *discordgo.Session, status string, activityType discordgo.ActivityType) {
	err := session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Status:    status,
		Activities: []*discordgo.Activity{{
			Name: activityOptions.Description,
			Type: activityType,
			URL:  activityOptions.Url,
		}},
	})

	sendErrMessage(err, "Cannot set the activity status!")
}

func sendErrMessage(value error, message string) {
	if value != nil {
		println(message)
	}
}
