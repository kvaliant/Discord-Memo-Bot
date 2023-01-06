package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
	"github.com/kvaliant/discord-memo-bot/helper"
	"github.com/kvaliant/discord-memo-bot/models"
)
func modalView(session *discordgo.Session, in *discordgo.InteractionCreate, memo *models.Memo) discordgo.InteractionResponse {
	// embed memoID in title Text Input CustomID
	titleTI := &discordgo.TextInput{}
	titleTI.CustomID = fmt.Sprintf("%v", memo.ID)
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
	data.Title = "View (and Update) Memo"
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
	return res
}

func (s *Service) MemoView(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	command := in.ApplicationCommandData()
	memoID := command.Options[0].Value.(float64)
	
	memo, err := controller.GetMemo(in.Member.User.ID, int(memoID))
	if err != nil { 
		res := helper.MemoNotFound(session, in)
		err = session.InteractionRespond(in.Interaction, &res)
		if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
		log.Printf("Err while fetching DB record, %v", err)
		return
	}
	
	isModalView := false
	if len(command.Options) > 1 {
		isModalView = command.Options[1].BoolValue()
	}

	var res discordgo.InteractionResponse
	if isModalView { 
		res = modalView(session, in, memo)
	} else {
		message := fmt.Sprintf("**%v** \n `%v`", memo.Title, memo.Content)

		data := &discordgo.InteractionResponseData{}
		data.Content = message

		res = discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: data,
		}
	}
	
	err = session.InteractionRespond(in.Interaction, &res)
	if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
}