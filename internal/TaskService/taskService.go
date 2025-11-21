package TaskService

type TaskService interface {
	CreateTask(task TaskStruct) (TaskStruct, error)
	GetAllTasks() ([]TaskStruct, error)
	GetTaskById(taskId uint) (TaskStruct, error)
	UpdateTask(taskId uint, task TaskStruct) (TaskStruct, error)
	DeleteTask(taskId uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (t *taskService) CreateTask(task TaskStruct) (TaskStruct, error) {
	newTask := TaskStruct{
		Title:     task.Title,
		Completed: task.Completed,
	}

	// Передаём указатель, чтобы GORM мог обновить поля
	if err := t.repo.CreateTask(&newTask); err != nil {
		return TaskStruct{}, err
	}

	// newTask теперь содержит ID, CreatedAt и UpdatedAt
	return newTask, nil
}

func (t *taskService) GetAllTasks() ([]TaskStruct, error) {
	return t.repo.GetAllTasks()
}

func (t *taskService) GetTaskById(taskId uint) (TaskStruct, error) {
	return t.repo.GetTaskById(taskId)
}

func (t *taskService) UpdateTask(taskId uint, task TaskStruct) (TaskStruct, error) {
	existingTask, err := t.repo.GetTaskById(taskId)
	if err != nil {
		return TaskStruct{}, err
	}

	existingTask.Title = task.Title
	existingTask.Completed = task.Completed

	if err := t.repo.UpdateTask(&existingTask); err != nil {
		return TaskStruct{}, err
	}

	return existingTask, nil
}

func (t *taskService) DeleteTask(taskId uint) error {
	return t.repo.DeleteTask(taskId)
}
