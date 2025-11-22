package handlers

import (
	"Basic/internal/TaskService"
	"Basic/internal/UsersService"
	"Basic/internal/web/users"
	"context"
)

type UserHandler struct {
	service     UsersService.UserService
	taskService TaskService.TaskService
}

func (u UserHandler) GetUser(ctx context.Context, request users.GetUserRequestObject) (users.GetUserResponseObject, error) {
	allusers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUser200JSONResponse{}
	for _, usr := range allusers {
		smuser := users.User{
			Id:        int64(usr.ID),
			Email:     usr.Email,
			Password:  usr.Password,
			CreatedAt: usr.CreatedAt,
			DeletedAt: *usr.DeletedAt, // Без проверок
			UpdatedAt: usr.UpdatedAt,
		}
		response = append(response, smuser)
	}
	return response, nil
}

func (u UserHandler) PostUser(ctx context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	userRequest := request.Body
	usertocreate := UsersService.UserStruct{
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}
	createduser, err := u.service.CreateUser(usertocreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUser201JSONResponse{
		Id:        int64(createduser.ID),
		Email:     createduser.Email,
		Password:  createduser.Password, // Показываем пароль
		CreatedAt: createduser.CreatedAt,
		UpdatedAt: createduser.UpdatedAt,
	}

	return response, nil
}

func (u UserHandler) DeleteUser(ctx context.Context, request users.DeleteUserRequestObject) (users.DeleteUserResponseObject, error) {
	id := uint(request.Id)
	err := u.service.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUser204Response{}
	return response, nil
}

func (u UserHandler) PatchUser(ctx context.Context, request users.PatchUserRequestObject) (users.PatchUserResponseObject, error) {
	id := uint(request.Id)
	updateinfo := request.Body

	usertoupdate := UsersService.UserStruct{
		Email:    *updateinfo.Email,
		Password: *updateinfo.Password,
	}

	updatedUser, err := u.service.UpdateUser(id, usertoupdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUser200JSONResponse{
		Id:        int64(updatedUser.ID),
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *UserHandler) GetTasksByUserId(ctx context.Context, request users.GetTasksByUserIdRequestObject) (users.GetTasksByUserIdResponseObject, error) {
	userId := uint(request.Id)

	// Получаем задачи через TaskService
	userTasks, err := u.taskService.GetAllTasksByUser(userId)
	if err != nil {
		return nil, err
	}

	// Преобразуем в формат API
	response := users.GetTasksByUserId200JSONResponse{}
	for _, tsk := range userTasks {
		task := users.Task{
			Id:        int64(tsk.ID),
			Title:     tsk.Title,
			Completed: tsk.Completed,
			UserId:    int64(tsk.UserID),
			CreatedAt: tsk.CreatedAt,
			UpdatedAt: tsk.UpdatedAt,
		}
		response = append(response, task)
	}

	return response, nil
}

func NewUserHandler(userService UsersService.UserService, taskService TaskService.TaskService) *UserHandler {
	return &UserHandler{
		taskService: taskService,
		service:     userService,
	}
}
