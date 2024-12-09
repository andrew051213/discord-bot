package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	ms := time.Now()

	if m.Author.Bot {
		return
	}

	elapsed := time.Since(ms) // e useless

	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Pong! Replied in %d ms", elapsed.Milliseconds()), m.Reference())
}
