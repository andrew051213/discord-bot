package functions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"time"
)

func BlacklistedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	bMessages := []string{"muie", "discord.gg", ".gg/", "steamcommunity", "https"}

	logChannel := os.Getenv("LOG_CHANNEL")

	for _, word := range bMessages {
		if strings.Contains(m.Content, word) {
			_, err := s.ChannelMessageSend(logChannel, "User <@"+m.Author.ID+"> has said "+word+" which is blacklisted")
			if err != nil {
				fmt.Printf("An error occured while attempting to send log message, %s", err)
			}

			timeoutDuration := time.Now().Add(1 * time.Hour) // 1 ora

			err = s.GuildMemberTimeout(m.GuildID, m.Author.ID, &timeoutDuration)

			if err != nil {
				fmt.Printf("An error occured while attempting to timeout user.")
			}

			s.ChannelMessageSend(logChannel, "User has been timeouted successfully")
		}
	}
}
