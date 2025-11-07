package main

import (
	"Basic/internal/TaskService"
	"Basic/internal/db"
	"Basic/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	TaskRepo := TaskService.NewTaskRepository(database)
	taskService := TaskService.NewTaskService(TaskRepo)
	TaskHandler := handlers.NewTaskHandler(taskService)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", TaskHandler.GetHandler)
	e.POST("/tasks", TaskHandler.PostHandler)
	e.PATCH("/tasks/:id", TaskHandler.PatchHandler)
	e.DELETE("/tasks/:id", TaskHandler.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
