package messagecreate

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/interface"
)

func Get() (eventHandler _interface.EventHandler) {
	eventHandler = _interface.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.MessageCreate); ok {
			println(event.Message.Content)
		}
	}}

	return eventHandler
}
