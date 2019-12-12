package main

import (
	"NepuBot/bot"
	"NepuBot/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Printf("an error occurred while reading the config, %s\n", err)
		return
	}

	bot.Run()
	//CTRL-C to terminate
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return
}