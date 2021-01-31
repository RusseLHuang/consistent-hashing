package entity

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primarykey";auto_increment;not_null`
	NationalID string `gorm:"unique"`
	Name       string
	Balance    uint `gorm:"default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time `sql:"DEFAULT:'current_timestamp'"`
	DeletedAt  time.Time `gorm:"index"`
}
