package TaskService

type Task struct {
	Task      string `json:"task"`
	ID        string `gorm:"primaryKey" json:"id"`
	Completed string `json:"completed"`
}
