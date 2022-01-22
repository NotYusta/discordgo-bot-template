package messagecreate

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/handler"
)

func Get() (eventHandler handler.EventHandler) {
	eventHandler = handler.EventHandler{Handle: func(session *discordgo.Session, event interface{}) {
		if event, ok := event.(*discordgo.MessageCreate); ok {
			println(event.Message.Content)
		}
	}}

	return eventHandler
}
