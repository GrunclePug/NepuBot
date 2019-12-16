package owner

import (
	"NepuBot/config"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func GuildList(s *discordgo.Session, m *discordgo.MessageCreate) {
	formattedGuilds := make([]string, len(s.State.Guilds))
	var guildList string

	userCount := 0
	for i := 0; i < len(s.State.Guilds); i++ {
		userCount += s.State.Guilds[i].MemberCount
	}

	for i := 0; i < len(s.State.Guilds); i++ {
		formattedGuilds[i] = "\nGuild: **" + s.State.Guilds[i].Name + "**\nUsers: " + strconv.Itoa(s.State.Guilds[i].MemberCount)
		guildList += formattedGuilds[i]
	}

	if m.Author.ID == config.Owner {
		embed := &discordgo.MessageEmbed{
			Title:       "NepuBot Guild List",
			Description: guildList,
			Color:       0xF9AF04,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Bot Info:",
					Value:  "Guilds: " + strconv.Itoa(len(s.State.Guilds)) + "\nUsers: " + strconv.Itoa(userCount),
					Inline: false,
				},
			},
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
