package commands

import (
	"auth/functions"
	"auth/modules"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	command := m.Content

	if strings.HasPrefix(command, "$") {
		command = strings.TrimPrefix(command, "$")

		args := strings.Fields(command)

		if len(args) < 1 {
			return
		}

		switch args[0] {
		case "ping":

			Ping(s, m)
		case "ban":
			if len(args) < 2 {
				modules.Reply(s, m, []string{"$ban user_id reason"})
			}

			modules.Ban(s, m, args[:2])
		case "unban":
			if len(args) < 2 {
				modules.Reply(s, m, []string{"$unban user_id"})
			}

			modules.UnBan(s, m, args[:2])

		default:
			modules.Reply(s, m, []string{"Command not found"})
		}
	}

	functions.BlacklistedMessage(s, m)
}
