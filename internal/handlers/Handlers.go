package handlers

import (
	"Basic/internal/TaskService"
	"Basic/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service TaskService.TaskService
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	alltasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range alltasks {
		smtask := tasks.Task{
			Id:        int64(tsk.ID),
			Title:     tsk.Title,
			Completed: tsk.Completed,
		}
		response = append(response, smtask)
	}
	return response, nil
}

func (h *TaskHandler) CreateTask(ctx context.Context, request tasks.CreateTaskRequestObject) (tasks.CreateTaskResponseObject, error) {
	taskRequest := request.Body
	tasktocreate := TaskService.Task{
		Title:     taskRequest.Title,
		Completed: *taskRequest.Completed,
	}
	createdtask, err := h.service.CreateTask(tasktocreate)
	if err != nil {
		return nil, err
	}
	response := tasks.CreateTask201JSONResponse{
		Id:        int64(createdtask.ID),
		Title:     createdtask.Title,
		Completed: createdtask.Completed,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, request tasks.DeleteTaskRequestObject) (tasks.DeleteTaskResponseObject, error) {
	id := uint(request.Id)
	err := h.service.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTask204Response{}
	return response, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, request tasks.UpdateTaskRequestObject) (tasks.UpdateTaskResponseObject, error) {
	id := uint(request.Id)

	// request.Body содержит поля для обновления
	updateinfo := request.Body

	// Получаем текущую задачу из БД
	tasktoupdate, err := h.service.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	if updateinfo.Title != nil {
		tasktoupdate.Title = *updateinfo.Title // разыменовываем указатель
	}
	if updateinfo.Completed != nil {
		tasktoupdate.Completed = *updateinfo.Completed
	}

	// Сохраняем обновленную задачу
	updatedTask, err := h.service.UpdateTask(id, tasktoupdate)
	if err != nil {
		return nil, err
	}
	response := tasks.UpdateTask200JSONResponse{
		Id:        int64(updatedTask.ID),
		Title:     updatedTask.Title,
		Completed: updatedTask.Completed,
		CreatedAt: updatedTask.CreatedAt,
		UpdatedAt: updatedTask.UpdatedAt,
	}
	return response, nil
}

func NewTaskHandler(service TaskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}
