package owner

import (
	"NepuBot/config"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func GuildInviteList(s *discordgo.Session, m *discordgo.MessageCreate) {
	formattedGuilds := make([]string, len(s.State.Guilds))
	var guildList string

	userCount := 0
	for i := 0; i < len(s.State.Guilds); i++ {
		userCount += s.State.Guilds[i].MemberCount
	}

	for i := 0; i < len(s.State.Guilds); i++ {
		var invList []*discordgo.Invite
		var inv = "not available"

		invList, _ = s.GuildInvites(s.State.Guilds[i].ID)
		if len(invList) > 0 {
			inv = invList[0].Code
		}
		formattedGuilds[i] = "\nGuild: **" + s.State.Guilds[i].Name + "**\nInvite: " + inv + "\nUsers: " + strconv.Itoa(s.State.Guilds[i].MemberCount)
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
		_ = s.ChannelTyping(m.ChannelID)
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
