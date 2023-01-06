package helper

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemoNotFound(session *discordgo.Session, in *discordgo.InteractionCreate) discordgo.InteractionResponse {	
	command := in.ApplicationCommandData()
	memoID := command.Options[0].Value.(float64)
	
	message := fmt.Sprintf("**ERR** memo ID `%v` not found/ not belong to current user", memoID)

	data := &discordgo.InteractionResponseData{}
	data.Content = message

	res := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	}

	return res
}