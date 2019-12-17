package general

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"time"
)

func ServerInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, _ := s.State.Guild(m.GuildID)
	owner, _ := s.State.Member(m.GuildID, guild.OwnerID)
	creationDate, _ := owner.JoinedAt.Parse()

	textChannels := 0
	voiceChannels := 0
	for _, channel := range guild.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannels++
		}
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			voiceChannels++
		}
	}

	embed := &discordgo.MessageEmbed{
		Title: "Server: " + guild.Name,
		Color: 0x00FF00,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    time.Now().Format(time.Stamp),
			IconURL: "https://i.imgur.com/WQSW5lV.png",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "・Owner",
				Value:  owner.Mention(),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・ID",
				Value:  guild.ID,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Region",
				Value:  strings.ReplaceAll(guild.Region, "-", " "),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Date Created",
				Value:  creationDate.Format(time.Stamp) + " (" + strconv.Itoa(int(time.Since(creationDate).Round(time.Hour).Hours()/24)) + " days ago)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Members",
				Value:  strconv.Itoa(guild.MemberCount),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Roles",
				Value:  strconv.Itoa(len(guild.Roles)),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Text Channels",
				Value:  strconv.Itoa(textChannels),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Voice Channels",
				Value:  strconv.Itoa(voiceChannels),
				Inline: false,
			},
		},
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
