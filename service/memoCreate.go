package service

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) MemoCreate(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	titleTI := &discordgo.TextInput{}
	titleTI.CustomID = "title"
	titleTI.Style = discordgo.TextInputShort
	titleTI.Label = "Title"
	titleTI.Placeholder = "Title"

	contentTI := &discordgo.TextInput{}
	contentTI.CustomID = "content"
	contentTI.Style = discordgo.TextInputParagraph
	contentTI.Label = "Content"
	contentTI.Placeholder = "Memo content"

	data := &discordgo.InteractionResponseData{}
	data.CustomID = "memo_create"
	data.Content = "nil"
	data.Title = "Create Memo"
	data.Components = []discordgo.MessageComponent {
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				*titleTI,
			},
		},
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				*contentTI,
			},
		},
	}

	res := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: data,
	}

	err := session.InteractionRespond(in.Interaction, &res)
	if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
}