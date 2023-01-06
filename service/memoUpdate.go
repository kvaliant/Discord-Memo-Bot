package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
	"github.com/kvaliant/discord-memo-bot/helper"
)

func (s *Service) MemoUpdate(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	command := in.ApplicationCommandData()
	memoID := command.Options[0].Value.(float64)

	memo, err := controller.GetMemo(in.Member.User.ID, int(memoID))
	if err != nil {
		if err != nil { 
			res := helper.MemoNotFound(session, in)
			err = session.InteractionRespond(in.Interaction, &res)
			if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
			log.Printf("Err while fetching DB record, %v", err)
			return
		}	
	}

	// embed memoID in title Text Input CustomID
	titleTI := &discordgo.TextInput{}
	titleTI.CustomID = fmt.Sprintf("%v", memoID)
	titleTI.Style = discordgo.TextInputShort
	titleTI.Label = "Title"
	titleTI.Placeholder = "Title"
	titleTI.Value = memo.Title	

	contentTI := &discordgo.TextInput{}
	contentTI.CustomID = "content"
	contentTI.Style = discordgo.TextInputParagraph
	contentTI.Label = "Content"
	contentTI.Placeholder = "Memo content"
	contentTI.Value = memo.Content

	data := &discordgo.InteractionResponseData{}
	data.CustomID = "memo_update"
	data.Content = "content"
	data.Title = "Update Memo"
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

	err = session.InteractionRespond(in.Interaction, &res)
	if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
}