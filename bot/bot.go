package bot

import (
	"NepuBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
	"time"
)

const Version = "v1.0"
var Session, _ = discordgo.New()

func Run() {
	var err error
	Session.Token = "Bot " + config.Token

	fmt.Printf(`
	 _      _____ ____  _     ____  ____  _____ 
	/ \  /|/  __//  __\/ \ /\/  __\/  _ \/__ __\
	| |\ |||  \  |  \/|| | ||| | //| / \|  / \  
	| | \|||  /_ |  __/| \_/|| |_\\| \_/|  | |  
	\_/  \|\____\\_/   \____/\____/\____/  \_/ %-16s`+"\n\n", Version)

	if Session.Token == "" {
		log.Println("You must provide a token!")
		return
	}

	Session.AddHandler(messageHandler)

	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	err = Session.UpdateListeningStatus(config.Prefix+"help | GrunclePug#7015")
	if err != nil {
		log.Printf("error setting discord status, %s\n", err)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, config.Prefix) {
		if m.Content == config.Prefix + "ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong! " + s.HeartbeatLatency().Round(time.Millisecond).String())
		}
	}
}