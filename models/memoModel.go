package models

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	ID				int			`gorm:"UNIQUE_INDEX:compositeindex" json:"id"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`

	DiscordUserID	string		`json:"-"`
}