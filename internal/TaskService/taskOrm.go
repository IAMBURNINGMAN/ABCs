package TaskService

import "time"

type TaskStruct struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Completed bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TaskStruct) TableName() string {
	return "tasks"
}
