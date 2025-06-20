package database

import "gorm.io/gorm"

type OTPModel struct {
	gorm.Model

	Secret string
}
