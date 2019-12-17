package general

import (
	"github.com/bwmarrin/discordgo"
)

func Invite(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Invite NepuBot to your server: <:nepubot:655682812904603658>",
		Description: "[Click here](https://discordapp.com/api/oauth2/authorize?client_id=430969608338669568&permissions=2146958583&scope=bot)",
		Color:       0x58FAF4,
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
