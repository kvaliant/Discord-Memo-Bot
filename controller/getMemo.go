package controller

import (
	"errors"

	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func GetMemo(userID string, memoID int) (*models.Memo, error) {
	var memo models.Memo
	res := config.DB.First(&memo, memoID)
	if res.Error != nil { return nil, res.Error }

	if memo.DiscordUserID != userID { return nil, errors.New("Unauthorized") }

	return &memo, nil
}