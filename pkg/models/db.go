package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migrate the schema
	return DB.AutoMigrate(&User{}, &Message{})
}
