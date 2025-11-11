package TaskService

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title" gorm:"type:text;not null"`
	Completed bool      `json:"completed" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Task) TableName() string {
	return "tasks"
}
