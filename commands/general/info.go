package general

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func Info(s *discordgo.Session, m *discordgo.MessageCreate) {
	userCount := 0
	for i := 0; i < len(s.State.Guilds); i++ {
		userCount += s.State.Guilds[i].MemberCount
	}

	embed := &discordgo.MessageEmbed{
		Title: "NepuBot Information",
		Color: 0xFF00FF,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Created by GrunclePug",
			IconURL: "https://i.imgur.com/mK2zlbr.png",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.imgur.com/73Cr7pY.png",
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "・About",
				Value:  "Multi-purpose open source Discord bot written in Go",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Description",
				Value:  "Multistage aerobic capacity test that progressively gets more difficult as it continues.",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Date Created",
				Value:  "Dec 11, 2019",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Guilds",
				Value:  strconv.Itoa(len(s.State.Guilds)),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Users",
				Value:  strconv.Itoa(userCount),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Website",
				Value:  "https://grunclepug.com",
				Inline: false,
			},
		},
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
