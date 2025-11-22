package TaskService

import "time"

type TaskStruct struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Completed bool      `gorm:"not null;default:false"`
	UserID    uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TaskStruct) TableName() string {
	return "tasks"
}
