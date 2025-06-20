package database

import "gorm.io/gorm"

type TelegramUserModel struct {
	gorm.Model

	UserID int64 `gorm:"primary_key;unique;not null;default:null;<-:create"`
}
