package controller

import (
	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func CreateMemo(userID string, title string, content string) (*models.Memo, error) {
	var memo models.Memo
	memo.Title = title
	memo.Content = content
	memo.DiscordUserID = userID

	res := config.DB.Create(&memo)
	if res.Error != nil {
		return nil, res.Error
	}

	return &memo, nil
}