package modules

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Ban(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		Reply(s, m, []string{"$ban id reason"})

		return
	}

	target_id := args[1]

	reason := strings.Join(args[2:], " ")

	if len(reason) == 0 {
		reason = "Not Specified"
	}

	isBanned := IsUserBanned(s, m.GuildID, target_id)

	if isBanned {
		Reply(s, m, []string{"User is already banned"})

		return
	}

	err := s.GuildBanCreateWithReason(m.GuildID, target_id, reason, 7)
	if err != nil {
		Reply(s, m, []string{"Couldn't ban user"})

		return
	}

	Reply(s, m, []string{"Successfully banned <@" + target_id + "> from the server."})
}
