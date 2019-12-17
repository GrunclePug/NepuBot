package general

//#include <time.h>
import "C"
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"runtime"
	"strconv"
	"time"
)

var startTime = time.Now()
var startTicks = C.clock()

func Resources(s *discordgo.Session, m *discordgo.MessageCreate) {
	clockSeconds := float64(C.clock()-startTicks) / float64(C.CLOCKS_PER_SEC)
	realSeconds := time.Since(startTime).Seconds()
	cpuUsage := clockSeconds / realSeconds * 100

	memory := &runtime.MemStats{}
	runtime.ReadMemStats(memory)

	embed := &discordgo.MessageEmbed{
		Title: "NepuBot System Info",
		Color: 0xFF00FF,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Created by GrunclePug",
			IconURL: "https://i.imgur.com/mK2zlbr.png",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: s.State.User.AvatarURL("512"),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "・CPU",
				Value:  "Usage: " + fmt.Sprintf("%.2f", cpuUsage) + "%" + "\nCores: " + strconv.Itoa(runtime.NumCPU()),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "・Memory",
				Value:  "Total: " + strconv.FormatUint(memory.Sys/1024000, 10) + "MB (" + fmt.Sprintf("%.1f", float64(memory.Sys/1024000)/1024) + "GB)",
				Inline: false,
			},
		},
	}
	_ = s.ChannelTyping(m.ChannelID)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
