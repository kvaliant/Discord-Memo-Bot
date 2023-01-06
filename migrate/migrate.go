package main

import (
	"github.com/kvaliant/discord-memo-bot/config"
	"github.com/kvaliant/discord-memo-bot/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(
		&models.Memo{},
	)
}
