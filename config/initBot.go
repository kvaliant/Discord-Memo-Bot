package config

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/structure"
)


func InitBot(cfg *structure.Config, callback func(session *discordgo.Session)) {
	session, err := discordgo.New("Bot "+cfg.Token)
	if err != nil { log.Fatalf("Error while initiating bot, %v", err) }

	session.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentGuildMessages
	callback(session)

	err = session.Open()
	if err != nil { log.Fatalf("Error while connecting to discord server, %v", err) }

	defer session.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Printf("Connection to discord server gracefully shutted down")
}