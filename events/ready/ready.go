package ready

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/interface"
)

func Get() (eventHandler _interface.EventHandler) {
	eventHandler = _interface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.Ready); ok {
			println("Logged in as " + event.User.Username)
		}
	}}

	return eventHandler
}
