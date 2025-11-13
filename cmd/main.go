package main

import (
	"Basic/internal/TaskService"
	"Basic/internal/db"
	"Basic/internal/handlers"
	"Basic/internal/web/tasks"
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

	strictHandler := tasks.NewStrictHandler(TaskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
