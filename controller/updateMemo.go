package controller

import (
	"errors"

	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func UpdateMemo(userID string, memoID int, title string, content string) (bool, error) {
	var memo models.Memo
	res := config.DB.First(&memo, memoID)
	if res.Error != nil { return false, res.Error }

	if memo.DiscordUserID != userID { return false, errors.New("Unauthorized") }

	if memo.Title == title && memo.Content == content {
		return false, nil
	}

	res = config.DB.Model(&memo).Updates(models.Memo{
		Title: title,
		Content: content,
	})

	if res.Error != nil { return false, res.Error }

	return true, nil
}