package commandinterface

import "github.com/bwmarrin/discordgo"

type CommandAccess string
type CommandsData map[string]*CommandInterface

const (
	CommandAccessDJ        CommandAccess = "DJ"
	CommandAccessAdmin     CommandAccess = "ADMIN"
	CommandAccessModerator CommandAccess = "MODERATOR"
)

type CommandInterface struct {
	Exec          func(session *discordgo.Session, event *discordgo.MessageCreate, args []string)
	Name          string
	Alias         []string
	CommandAccess []CommandAccess
}
