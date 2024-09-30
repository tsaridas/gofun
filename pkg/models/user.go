package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;not null" json:"username"`
	Email    string    `gorm:"uniqueIndex;not null" json:"email"`
	Messages []Message `gorm:"foreignKey:UserID" json:"messages,omitempty"`
}

func (u *User) FullName() string {
	return u.Username
}

func GetUser(id uint) (*User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}
