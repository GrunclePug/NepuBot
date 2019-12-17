package general

import (
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "NepuBot Command Help",
		Description: "Bot created by @GrunclePug#7015",
		Color:       0xFF00FF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.imgur.com/73Cr7pY.png",
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Commands",
				Value: "**help** | Brings you here" +
					"\n**info** | See info on the bot" +
					"\n**ping** | Checks the delay between you and the bot" +
					"\n**invite** | gives you the link to invite the bot",
				Inline: false,
			},
		},
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
