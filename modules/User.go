package modules

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func IsUserBanned(s *discordgo.Session, guildID, userID string) bool {
	bans, err := s.GuildBans(guildID, 1000, "", "")
	if err != nil {
		fmt.Println("Error fetching bans:", err)

		return false
	}

	for _, ban := range bans {
		if ban.User.ID == userID {
			return true
		}
	}

	return false
}
