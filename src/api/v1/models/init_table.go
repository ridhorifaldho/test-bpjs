package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&Order{})
}
