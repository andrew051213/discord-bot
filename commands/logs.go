package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func Logs(s *discordgo.Session, l *discordgo.GuildRoleCreate) {
	fmt.Println("A role has been created")

	channelID := os.Getenv("LOG_CHANNEL")

	_, err := s.ChannelMessageSend(channelID, "A new role has been created <@&"+l.Role.ID+">")
	if err != nil {
		fmt.Printf("An error occured while attempting to send logs, %s", err)

		return
	}
}
