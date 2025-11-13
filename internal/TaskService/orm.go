package TaskService

import "time"

type Task struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Completed bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Task) TableName() string {
	return "tasks"
}
