package handlers

import (
	"Basic/internal/UsersService"
	"Basic/internal/web/users"
	"context"
)

type UserHandler struct {
	service UsersService.UserService
}

func (u UserHandler) GetUser(ctx context.Context, request users.GetUserRequestObject) (users.GetUserResponseObject, error) {
	allusers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUser200JSONResponse{}
	for _, tsk := range allusers {
		smuser := users.User{
			Id:        int64(tsk.ID),
			Email:     tsk.Email,
			Password:  tsk.Password,
			CreatedAt: tsk.CreatedAt,
			DeletedAt: tsk.DeletedAt,
			UpdatedAt: tsk.UpdatedAt,
		}
		response = append(response, smuser)
	}
	return response, nil
}

func (u UserHandler) PostUser(ctx context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	userRequest := request.Body
	usertocreate := UsersService.UserStruct{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	createduser, err := u.service.CreateUser(usertocreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUser201JSONResponse{
		Id:        int64(createduser.ID),
		Email:     createduser.Email,
		Password:  createduser.Password,
		CreatedAt: createduser.CreatedAt,
		DeletedAt: createduser.DeletedAt,
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

	usertoupdate, err := u.service.GetUserById(id)
	if err != nil {
		return nil, err
	}

	usertoupdate.Email = updateinfo.Email
	usertoupdate.Password = updateinfo.Password

	updatedUser, err := u.service.UpdateUser(id, usertoupdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUser200JSONResponse{
		Id:        int64(updatedUser.ID),
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
		CreatedAt: updatedUser.CreatedAt,
		DeletedAt: updatedUser.DeletedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}
	return response, nil
}

func NewUserHandler(service UsersService.UserService) *UserHandler {
	return &UserHandler{service: service}
}
