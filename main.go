package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	Task      string `json:"task"`
	ID        string `gorm:"primaryKey" json:"id"`
	Completed string `json:"completed"`
}

var (
	db *gorm.DB
	mu sync.RWMutex
)

func initDB() {

	dsn := "host=localhost user=postgres password=eourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
}

func GetHandler(c echo.Context) error {
	tasks := []Task{}

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get tasks from database"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostHandler(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bind error"})
	}
	mu.Lock()
	task := Task{req.Task, uuid.NewString(), req.Completed}
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task to database"})
	}
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

	var task Task
	if err := db.First(&task, "id = ?", idpatcher).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find  this task to patch"})
	}

	task.Task = req.Task
	task.Completed = req.Completed
	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to patch task to database"})
	}
	return c.JSON(http.StatusOK, task)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task from database"})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", GetHandler)
	e.POST("/tasks", PostHandler)
	e.PATCH("/tasks/:id", PatchHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
