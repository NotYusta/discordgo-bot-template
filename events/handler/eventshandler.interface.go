package handler

import "github.com/bwmarrin/discordgo"

type EventHandler struct {
	Handle func(session *discordgo.Session, event interface{})
}
