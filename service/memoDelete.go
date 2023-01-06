package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
	"github.com/kvaliant/discord-memo-bot/helper"
)

func (s *Service) MemoDelete(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	command := in.ApplicationCommandData()
	memoID := command.Options[0].Value.(float64)
	
	err := controller.DeleteMemo(in.Member.User.ID, int(memoID))
	if err != nil { 
		if err != nil { 
			res := helper.MemoNotFound(session, in)
			err = session.InteractionRespond(in.Interaction, &res)
			if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
			log.Printf("Err while fetching DB record, %v", err)
			return
		}	
	}
	
	message := fmt.Sprintf("Deleted : 1")

	data := &discordgo.InteractionResponseData{}
	data.Content = message

	res := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	}
	
	err = session.InteractionRespond(in.Interaction, &res)
	if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
}