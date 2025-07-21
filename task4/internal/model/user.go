package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:32"`
	Password string `gorm:"size:255"`
	Role     string `gorm:"size:32;default:user"`
}
