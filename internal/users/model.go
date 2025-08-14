package users

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex:idx_user_name"`
	DisplayName string `gorm:"uniqueIndex:idx_display_name"`
	Email       string `gorm:"uniqueIndex:idx_email"`
	Token       string `gorm:"unique"`
	Hash        string
	Color       string
}
