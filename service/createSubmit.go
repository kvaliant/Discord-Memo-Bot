package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/discord-memo-bot/controller"
)

func (s *Service) CreateSubmit(session *discordgo.Session, in *discordgo.InteractionCreate) {	
	mc := in.Interaction.ModalSubmitData().Components[0]
	byteJson, _ := mc.MarshalJSON()
	var ar discordgo.ActionsRow
	if err := ar.UnmarshalJSON(byteJson); err != nil { log.Panicf("Err while unmarshaling action row component, %v", err) }

	mc = ar.Components[0]
	byteJson, _ = mc.MarshalJSON()
	var titleTI discordgo.TextInput
	if err := discordgo.Unmarshal(byteJson, &titleTI); err != nil { log.Panicf("Err while unmarshaling text input component, %v", err) }


	mc = in.Interaction.ModalSubmitData().Components[1]
	byteJson, _ = mc.MarshalJSON()
	var ar2 discordgo.ActionsRow
	if err := ar2.UnmarshalJSON(byteJson); err != nil { log.Panicf("Err while unmarshaling action row component, %v", err) }

	mc = ar2.Components[0]
	byteJson, _ = mc.MarshalJSON()
	var contentTI discordgo.TextInput
	if err := discordgo.Unmarshal(byteJson, &contentTI); err != nil { log.Panicf("Err while unmarshaling text input component, %v", err) }
	
	memo, err := controller.CreateMemo(in.Member.User.ID, titleTI.Value, contentTI.Value)
	if err != nil { log.Panicf("Err while creating DB records, %v", err) }

	content := fmt.Sprintf("Memo created at ID:%v", memo.ID)
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	err = session.InteractionRespond(in.Interaction, &data)
	if err != nil { log.Panicf("Err while interaction respond, %v", err) }
}