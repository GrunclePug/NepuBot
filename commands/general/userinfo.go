package general

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func UserInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	var user = m.Author
	var err error

	if len(m.Mentions) > 0 {
		user = m.Mentions[0]
	}

	guild, _ := s.State.Guild(m.GuildID)
	mem, _ := s.State.Member(guild.ID, user.ID)
	joinDate, _ := mem.JoinedAt.Parse()
	presence, _ := s.State.Presence(guild.ID, user.ID)

	status := "unknown"
	switch presence.Status {
	case discordgo.StatusOnline:
		status = "Online <:online:656053599163645974>"
	case discordgo.StatusDoNotDisturb:
		status = "Do Not Disturb <:dnd:656053616544710657>"
	case discordgo.StatusIdle:
		status = "Idle <:idle:656053641840427008>"
	case discordgo.StatusOffline:
		status = "Offline <:offline:656053661230825473>"
	}

	nickname := "n/a"
	if mem.Nick != "" {
		nickname = mem.Nick
	}

	embed := &discordgo.MessageEmbed{
		Title: "User Info:",
		Color: 0x00FF00,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    time.Now().Format(time.Stamp),
			IconURL: "https://i.imgur.com/WQSW5lV.png",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL("512"),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "・User",
				Value:  user.Username + "#" + user.Discriminator,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Nickname",
				Value:  nickname,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・ID",
				Value:  user.ID,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Status",
				Value:  status,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Account Created",
				Value:  "coming soon",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Server Join Date",
				Value:  joinDate.Format(time.Stamp) + " (" + strconv.Itoa(int(time.Since(joinDate).Round(time.Hour).Hours()/24)) + " days ago)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Roles",
				Value:  strconv.Itoa(len(mem.Roles)),
				Inline: false,
			},
		},
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	fmt.Println(err)
}
