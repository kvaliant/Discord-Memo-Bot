package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
	"github.com/kvaliant/discord-memo-bot/models"
)

func (s *Service) ListAll(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	var memos []models.Memo

	memos, err := controller.GetMemos(in.Member.User.ID)
	if err != nil { log.Panicf("Err while fetching DB records, %v", err) }

	message := "**Here is the list of all your memo : **"
	for _, memo := range memos{
		message = fmt.Sprintf("%v \n %v \t | \t `%v`", message, memo.ID, memo.Title)
	}

	data := &discordgo.InteractionResponseData{}
	data.Content = message

	res := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	}
	
	err = session.InteractionRespond(in.Interaction, &res)
	if err != nil { log.Panicf("Err while responding to interaction, %v", err) }
}