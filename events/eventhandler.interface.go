package events

import "github.com/bwmarrin/discordgo"

type EventHandler struct {
	Handler func(session discordgo.Session, event interface{})
}
