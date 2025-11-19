package TaskService

import "github.com/stretchr/testify/mock"

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *TaskStruct) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks() ([]TaskStruct, error) {
	args := m.Called()
	if tasks, ok := args.Get(0).([]TaskStruct); ok {
		return tasks, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) GetTaskById(id uint) (TaskStruct, error) {
	args := m.Called(id)
	return args.Get(0).(TaskStruct), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task *TaskStruct) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
