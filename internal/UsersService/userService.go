package UsersService

type UserService interface {
	CreateUser(user UserStruct) (UserStruct, error)
	GetAllUsers() ([]UserStruct, error)
	GetUserById(id uint) (UserStruct, error)
	UpdateUser(id uint, user UserStruct) (UserStruct, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo UserRepository
}

func (u userService) CreateUser(user UserStruct) (UserStruct, error) {
	newuser := UserStruct{
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.repo.CreateUser(&newuser); err != nil {
		return UserStruct{}, err
	}
	return newuser, nil
}

func (u userService) GetAllUsers() ([]UserStruct, error) {
	return u.repo.GetAllUsers()
}

func (u userService) GetUserById(id uint) (UserStruct, error) {
	return u.repo.GetUserById(id)
}

func (u userService) UpdateUser(id uint, user UserStruct) (UserStruct, error) {
	existingUser, err := u.repo.GetUserById(id)
	if err != nil {
		return UserStruct{}, err
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if err := u.repo.UpdateUser(&existingUser); err != nil {
		return UserStruct{}, err
	}
	return existingUser, nil
}

func (u userService) DeleteUser(id uint) error {
	return u.repo.DeleteUser(id)
}

func NewUserService(r UserRepository) UserService { return &userService{repo: r} }
