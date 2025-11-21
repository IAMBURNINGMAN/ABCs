package TaskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task *TaskStruct) error
	GetAllTasks() ([]TaskStruct, error)
	GetTaskById(taskId uint) (TaskStruct, error)
	UpdateTask(task *TaskStruct) error
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *TaskStruct) error {
	// GORM сам заполнит ID и CreatedAt / UpdatedAt
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAllTasks() ([]TaskStruct, error) {
	var tasks []TaskStruct
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskById(taskId uint) (TaskStruct, error) {
	var task TaskStruct
	err := r.db.First(&task, taskId).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task *TaskStruct) error {
	// GORM обновит UpdatedAt автоматически
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(taskId uint) error {
	return r.db.Delete(&TaskStruct{}, taskId).Error
}
