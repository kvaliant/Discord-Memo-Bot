package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/handler"
	"github.com/kvaliant/discord-memo-bot/service"
	"github.com/kvaliant/discord-memo-bot/structure"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	cfg := structure.Config{}
	cfg.Token = os.Getenv("TOKEN")
	cfg.Application = os.Getenv("APPLICATION")

	log.Printf("Starting Discord Memo Bot Application\n")
	log.Printf("Configuration:\n")
	log.Printf("TOKEN : %v\n", cfg.Token)
	log.Printf("APPLICATION ID : %v\n", cfg.Application)

	config.InitBot(&cfg, func(session *discordgo.Session){
		log.Printf("Discord server connection established")

		service := service.NewBotService()
		h := handler.NewBotHandler(service)
		session.AddHandler(h.OnReady)
		session.AddHandler(h.OnInteraction)
	})
}

