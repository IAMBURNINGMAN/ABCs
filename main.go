package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TaskReq struct {
	Task string `json:"task"`
}

type TaskRes struct {
	Task string `json:"task"`
}

var (
	lastTask string
	mu       sync.RWMutex
)

func PostHandler(c echo.Context) error {
	var req TaskReq
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bind error"})
	}
	mu.Lock()
	lastTask = req.Task
	mu.Unlock()
	return c.JSON(http.StatusOK, TaskRes{Task: lastTask})
}

func GetHandler(c echo.Context) error {
	mu.RLock()
	t := lastTask
	mu.RUnlock()
	if t == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no task set"})
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("привет, %s", t))
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
