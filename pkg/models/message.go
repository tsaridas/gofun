package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint   `gorm:"not null" json:"user_id"`
	Content string `gorm:"not null" json:"content"`
	User    User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func CreateMessage(message *Message) error {
	return DB.Create(message).Error
}

func GetMessagesForUser(userID uint) ([]Message, error) {
	var messages []Message
	result := DB.Where("user_id = ?", userID).Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
