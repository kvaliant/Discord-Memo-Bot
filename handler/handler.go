package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/service"
)

type handler struct {
	service service.Service
}

func NewBotHandler(svc service.Service) handler {
	return handler{service: svc}
}

func (h *handler) OnReady(session *discordgo.Session, message *discordgo.Ready) {
	log.Print("Discord bot active, Listening for events")
}

func (h *handler) OnInteraction(session *discordgo.Session, in *discordgo.InteractionCreate) {
	if in.Type == discordgo.InteractionApplicationCommand {
		switch commandName := in.ApplicationCommandData().Name; commandName {
		case "memo_list":
			h.service.ListAll(session, in)
		case "memo_view":
			h.service.MemoView(session, in)
		case "memo_create":
			h.service.MemoCreate(session, in)
		case "memo_update":
			h.service.MemoUpdate(session, in)
		case "memo_delete":
			h.service.MemoDelete(session, in)
		}

	}

	if in.Type == discordgo.InteractionModalSubmit {
		switch modalID := in.ModalSubmitData().CustomID; modalID {
		case "memo_create":
			h.service.CreateSubmit(session, in)
		case "memo_update":
			h.service.UpdateSubmit(session, in)
		}
	}
	// h.service.InteractionNotFound(session, in)
}