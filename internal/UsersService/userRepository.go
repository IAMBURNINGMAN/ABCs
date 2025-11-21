package UsersService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *UserStruct) error
	GetAllUsers() ([]UserStruct, error)
	GetUserById(id uint) (UserStruct, error)
	UpdateUser(user *UserStruct) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) CreateUser(user *UserStruct) error {
	return u.db.Create(user).Error
}

func (u userRepository) GetAllUsers() ([]UserStruct, error) {
	var users []UserStruct
	err := u.db.Find(&users).Error
	return users, err
}

func (u userRepository) GetUserById(id uint) (UserStruct, error) {
	var thatUser UserStruct
	err := u.db.First(&thatUser, id).Error
	return thatUser, err
}

func (u userRepository) UpdateUser(user *UserStruct) error {
	return u.db.Save(user).Error
}

func (u userRepository) DeleteUser(id uint) error {
	return u.db.Delete(&UserStruct{}, id).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
