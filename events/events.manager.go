package events

import "github.com/bwmarrin/discordgo"

var (
	eventHandlers = []EventHandler{}
)

func EventManager(session discordgo.Session) {
	for _, element := range eventHandlers {
		session.AddHandler(element.Handle)
	}
}
