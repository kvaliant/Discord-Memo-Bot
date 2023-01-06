package controller

import (
	"errors"

	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func DeleteMemo(userID string, memoID int) (error) {
	var memo models.Memo
	res := config.DB.First(&memo, memoID)
	if res.Error != nil { return res.Error }

	if memo.DiscordUserID != userID { return errors.New("Unauthorized") }

	res = config.DB.Delete(&models.Memo{}, memoID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}