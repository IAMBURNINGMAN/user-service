package user

import (
	"time"
)

type UserStruct struct {
	ID        uint       `gorm:"primaryKey"`
	Email     string     `gorm:"not null;uniqueIndex"`
	Password  string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (UserStruct) TableName() string {
	return "users"
}
