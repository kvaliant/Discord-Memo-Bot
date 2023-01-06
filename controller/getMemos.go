package controller

import (
	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func GetMemos(userID string) ([]models.Memo, error) {
	var memos []models.Memo
	res := config.DB.Where("discord_user_id = ?", userID).Find(&memos)

	if res.Error != nil { return nil, res.Error }

	return memos, nil
}