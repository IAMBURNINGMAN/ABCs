package TaskService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	tests := []struct {
		name      string
		input     TaskStruct
		mockSetup func(m *MockTaskRepository, input *TaskStruct)
		wantErr   bool
	}{
		{
			name:  "successful creation",
			input: TaskStruct{Title: "Test", Completed: false},
			mockSetup: func(m *MockTaskRepository, input *TaskStruct) {
				m.On("CreateTask", input).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "db error",
			input: TaskStruct{Title: "Bad test", Completed: false},
			mockSetup: func(m *MockTaskRepository, input *TaskStruct) {
				m.On("CreateTask", input).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)

			tt.mockSetup(mockRepo, &tt.input)

			service := NewTaskService(mockRepo)

			result, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllTasks(t *testing.T) {

	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		want      []TaskStruct
		wantErr   bool
	}{
		{
			name: "successful - many tasks",
			mockSetup: func(m *MockTaskRepository) {
				tasks := []TaskStruct{
					{Title: "t1", Completed: false},
					{Title: "t2", Completed: false},
				}
				m.On("GetAllTasks").Return(tasks, nil)
			},
			want: []TaskStruct{
				{Title: "t1", Completed: false},
				{Title: "t2", Completed: false},
			},
			wantErr: false,
		},
		{
			name: "empty result",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]TaskStruct{}, nil)
			},
			want:    []TaskStruct{},
			wantErr: false,
		},
		{
			name: "db error",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return(nil, errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)

			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)

			result, err := service.GetAllTasks()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		id        uint
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   1,
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", uint(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "db error",
			id:   2,
			mockSetup: func(m *MockTaskRepository) {
				m.On("DeleteTask", uint(2)).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)
			service := NewTaskService(mockRepo)
			err := service.DeleteTask(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
func TestGetTaskById(t *testing.T) {

	tests := []struct {
		name      string
		id        uint
		mockSetup func(m *MockTaskRepository)
		want      TaskStruct
		wantErr   bool
	}{
		{
			name: "success",
			id:   1,
			mockSetup: func(m *MockTaskRepository) {
				task := TaskStruct{ID: 1, Title: "Test", Completed: false}
				m.On("GetTaskById", uint(1)).Return(task, nil)
			},
			want:    TaskStruct{ID: 1, Title: "Test", Completed: false},
			wantErr: false,
		},
		{
			name: "not found",
			id:   2,
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", uint(2)).Return(TaskStruct{}, errors.New("not found"))
			},
			want:    TaskStruct{},
			wantErr: true,
		},
		{
			name: "db error",
			id:   3,
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTaskById", uint(3)).Return(TaskStruct{}, errors.New("db error"))
			},
			want:    TaskStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := new(MockTaskRepository)

			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)

			result, err := service.GetTaskById(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {

	tests := []struct {
		name      string
		taskId    uint
		input     TaskStruct
		mockSetup func(m *MockTaskRepository, taskId uint, input TaskStruct)
		want      TaskStruct
		wantErr   bool
	}{
		{
			name:   "successful update",
			taskId: 1,
			input: TaskStruct{
				Title:     "Updated title",
				Completed: true,
			},
			mockSetup: func(m *MockTaskRepository, taskId uint, input TaskStruct) {

				existing := TaskStruct{
					ID:        taskId,
					Title:     "Old title",
					Completed: false,
				}

				m.On("GetTaskById", taskId).Return(existing, nil)

				updated := TaskStruct{
					ID:        taskId,
					Title:     input.Title,
					Completed: input.Completed,
				}

				m.On("UpdateTask", &updated).Return(nil)
			},
			want: TaskStruct{
				ID:        1,
				Title:     "Updated title",
				Completed: true,
			},
			wantErr: false,
		},

		{
			name:   "GetTaskById returns error",
			taskId: 2,
			input: TaskStruct{
				Title:     "New title",
				Completed: true,
			},
			mockSetup: func(m *MockTaskRepository, taskId uint, input TaskStruct) {
				m.On("GetTaskById", taskId).Return(TaskStruct{}, errors.New("not found"))
			},
			want:    TaskStruct{},
			wantErr: true,
		},

		{
			name:   "UpdateTask returns error",
			taskId: 3,
			input: TaskStruct{
				Title:     "Will fail",
				Completed: false,
			},
			mockSetup: func(m *MockTaskRepository, taskId uint, input TaskStruct) {

				existing := TaskStruct{
					ID:        taskId,
					Title:     "Old",
					Completed: false,
				}

				m.On("GetTaskById", taskId).Return(existing, nil)

				updated := TaskStruct{
					ID:        taskId,
					Title:     input.Title,
					Completed: input.Completed,
				}

				m.On("UpdateTask", &updated).Return(errors.New("db error"))
			},
			want:    TaskStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.taskId, tt.input)

			service := NewTaskService(mockRepo)
			result, err := service.UpdateTask(tt.taskId, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
