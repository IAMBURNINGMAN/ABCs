package handlers

import (
	"Basic/internal/TaskService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service TaskService.TaskService
}

func NewTaskHandler(service TaskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not get tasks from db"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var req TaskService.Task
	tasking, err := h.service.CreateTask(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create task"})
	}
	return c.JSON(http.StatusOK, tasking)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	idpatcher := c.Param("id")
	var req TaskService.Task

	updatedtask, err := h.service.UpdateTask(idpatcher, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
	}
	return c.JSON(http.StatusOK, updatedtask)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete task"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
