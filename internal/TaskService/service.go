package TaskService

import (
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(taskId string) (Task, error)
	UpdateTask(taskId string, task Task) (Task, error)
	DeleteTask(taskId string) error
}
type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (t *taskService) CreateTask(task Task) (Task, error) {
	tasking := Task{
		Task:      task.Task,
		ID:        uuid.NewString(),
		Completed: task.Completed,
	}
	if err := t.repo.CreateTask(tasking); err != nil {
		return Task{}, err
	}
	return tasking, nil
}

func (t *taskService) GetAllTasks() ([]Task, error) {
	return t.repo.GetAllTasks()
}

func (t *taskService) GetTaskById(taskId string) (Task, error) {
	return t.repo.GetTaskById(taskId)
}

func (t *taskService) UpdateTask(taskId string, task Task) (Task, error) {
	tasking, err := t.repo.GetTaskById(taskId)
	if err != nil {
		return Task{}, err
	}
	tasking.Task = task.Task
	tasking.Completed = task.Completed
	if err := t.repo.UpdateTask(tasking); err != nil {
		return Task{}, err
	}
	return tasking, nil
}

func (t *taskService) DeleteTask(taskId string) error {
	return t.repo.DeleteTask(taskId)
}
