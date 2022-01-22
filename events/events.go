package events

import (
	"github.com/bwmarrin/discordgo"
	"go-dc-bot/events/handler"
	"go-dc-bot/events/messagecreate"
	"go-dc-bot/events/ready"
)

var (
	EventHandlers = []handler.EventHandler{
		messagecreate.Get(),
		ready.Get(),
	}
)

func Init(session *discordgo.Session) {
	for _, element := range EventHandlers {
		session.AddHandler(element.Handle)
	}
}
