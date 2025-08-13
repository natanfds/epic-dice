package rooms

import "gorm.io/gorm"

type RoomModel struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
}
