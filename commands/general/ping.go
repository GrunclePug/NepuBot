package general

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Pong! :ping_pong:",
		Description: "pong! " + s.HeartbeatLatency().Round(time.Millisecond).String(),
		Color:       0xFE2E2E,
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
