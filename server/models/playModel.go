package models

import "gorm.io/gorm"

type Play struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	CreatorID   uint   `gorm:"not null" json:"creatorId"`
}
