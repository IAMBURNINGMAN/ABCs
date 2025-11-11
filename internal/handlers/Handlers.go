package handlers

import (
	"Basic/internal/TaskService"
	"net/http"
	"strconv"

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
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title is required"})
	}

	created, err := h.service.CreateTask(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create task"})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
	}

	var req TaskService.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	updated, err := h.service.UpdateTask(uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update task"})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	// Получаем ID из URL параметра
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
	}

	if err := h.service.DeleteTask(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete task"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
