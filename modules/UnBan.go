package modules

import (
	"github.com/bwmarrin/discordgo"
)

func UnBan(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 1 {
		Reply(s, m, []string{"$unban user_id"})

		return
	}

	target_id := args[1]

	if len(target_id) == 0 {
		Reply(s, m, []string{"Please specify a user_id to unban"})

		return
	}

	isBanned := IsUserBanned(s, m.GuildID, target_id)

	if !isBanned {
		Reply(s, m, []string{"User is not banned"})

		return
	}

	err := s.GuildBanDelete(m.GuildID, target_id)
	if err != nil {
		Reply(s, m, []string{"An error occured while attempting to unban <@" + target_id + ">. Please try again later"})

		return
	}

	Reply(s, m, []string{"Successfully unbanned user <@" + target_id + ">."})
}
