package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
)

func (s *Service) UpdateSubmit(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	mc := in.Interaction.ModalSubmitData().Components[0]
	byteJson, _ := mc.MarshalJSON()
	var ar discordgo.ActionsRow
	if err := ar.UnmarshalJSON(byteJson); err != nil { log.Panicf("Err while unmarshaling action row component, %v", err) }

	mc = ar.Components[0]
	byteJson, _ = mc.MarshalJSON()
	var titleTI discordgo.TextInput
	if err := discordgo.Unmarshal(byteJson, &titleTI); err != nil { log.Panicf("Err while unmarshaling text input component, %v", err) }

	memoID, _ := strconv.Atoi(titleTI.CustomID)

	mc = in.Interaction.ModalSubmitData().Components[1]
	byteJson, _ = mc.MarshalJSON()
	var ar2 discordgo.ActionsRow
	if err := ar2.UnmarshalJSON(byteJson); err != nil { log.Panicf("Err while unmarshaling action row component, %v", err) }

	mc = ar2.Components[0]
	byteJson, _ = mc.MarshalJSON()
	var contentTI discordgo.TextInput
	if err := discordgo.Unmarshal(byteJson, &contentTI); err != nil { log.Panicf("Err while unmarshaling text input component, %v", err) }

	updated, err := controller.UpdateMemo(in.Member.User.ID, memoID, titleTI.Value, contentTI.Value)
	if err != nil { log.Panicf("Err while creating DB records, %v", err) }

	var content string 
	if updated {
		content = fmt.Sprintf("Memo updated")
	} else {
		content = fmt.Sprintf("Memo not updated")
	}
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	err = session.InteractionRespond(in.Interaction, &data)
	if err != nil { log.Panicf("Err while interaction respond, %v", err) }
}