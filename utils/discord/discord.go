package discord

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

//goland:noinspection GoUnhandledErrorResult,GoUnusedCallResult
func ChannelMessageSend(session *discordgo.Session, channelId string, message string) *discordgo.Message {
	send, err := session.ChannelMessageSend(channelId, message)

	if err == nil {
		return send
	}

	errors.New(err.Error())
	return nil
}
