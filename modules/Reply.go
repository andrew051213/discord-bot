package modules

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Reply(s *discordgo.Session, m *discordgo.MessageCreate, message []string) {
	if len(message) == 0 {
		return
	}

	reply_message := strings.Join(message, " ")

	s.ChannelMessageSendReply(m.ChannelID, reply_message, m.Reference())
}
