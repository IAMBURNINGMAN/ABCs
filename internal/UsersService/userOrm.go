package UsersService

import (
	"time"
)

type UserStruct struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"not null;uniqueIndex"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time `gorm:"autoDeleteTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (UserStruct) TableName() string {
	return "users"
}
