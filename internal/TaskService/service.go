package TaskService

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(taskId uint) (Task, error)
	UpdateTask(taskId uint, task Task) (Task, error)
	DeleteTask(taskId uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (t *taskService) CreateTask(task Task) (Task, error) {
	// ID будет сгенерирован автоматически базой данных (SERIAL)
	newTask := Task{
		Title:     task.Title,
		Completed: task.Completed,
	}

	if err := t.repo.CreateTask(newTask); err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (t *taskService) GetAllTasks() ([]Task, error) {
	return t.repo.GetAllTasks()
}

func (t *taskService) GetTaskById(taskId uint) (Task, error) {
	return t.repo.GetTaskById(taskId)
}

func (t *taskService) UpdateTask(taskId uint, task Task) (Task, error) {
	existingTask, err := t.repo.GetTaskById(taskId)
	if err != nil {
		return Task{}, err
	}

	existingTask.Title = task.Title
	existingTask.Completed = task.Completed

	if err := t.repo.UpdateTask(existingTask); err != nil {
		return Task{}, err
	}
	return existingTask, nil
}

func (t *taskService) DeleteTask(taskId uint) error {
	return t.repo.DeleteTask(taskId)
}
