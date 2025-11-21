package main

import (
	"Basic/internal/TaskService"
	"Basic/internal/UsersService"
	"Basic/internal/db"
	"Basic/internal/handlers"
	"Basic/internal/web/tasks"
	"Basic/internal/web/users"
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
	TaskServicee := TaskService.NewTaskService(TaskRepo)
	TaskHandler := handlers.NewTaskHandler(TaskServicee)
	UserRepo := UsersService.NewUserRepository(database)
	UserService := UsersService.NewUserService(UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(TaskHandler, nil)
	strictHandlerUS := users.NewStrictHandler(UserHandler, nil)

	tasks.RegisterHandlers(e, strictHandler)
	users.RegisterHandlers(e, strictHandlerUS)

	e.Logger.Fatal(e.Start(":8080"))
}
