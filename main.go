package main

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	Task      string `json:"task"`
	ID        string `json:"id"`
	Completed string `json:"completed"`
}

var (
	mu sync.RWMutex
)

var TASKLIST = []Task{}

func PostHandler(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bind error"})
	}
	mu.Lock()
	task := Task{req.Task, uuid.NewString(), req.Completed}
	TASKLIST = append(TASKLIST, task)
	mu.Unlock()
	return c.JSON(http.StatusOK, Task{Task: req.Task, ID: req.ID, Completed: req.Completed})
}

func PatchHandler(c echo.Context) error {
	idpatcher := c.Param("id")
	var req Task
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bind error"})
	}
	for i, tasking := range TASKLIST {
		if tasking.ID == idpatcher {
			TASKLIST[i].Completed = req.Completed
			TASKLIST[i].Task = req.Task
			return c.JSON(http.StatusOK, TASKLIST[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func GetHandler(c echo.Context) error {
	if len(TASKLIST) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no task set"})
	}
	mu.RLock()
	t := TASKLIST[0]
	mu.RUnlock()
	if len(TASKLIST) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no task set"})
	}
	return c.JSON(http.StatusOK, t)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")
	for i, t := range TASKLIST {
		if t.ID == id {
			TASKLIST = append(TASKLIST[:i], TASKLIST[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)
	e.PATCH("/task/:id", PatchHandler)
	e.DELETE("/task/:id", deleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
