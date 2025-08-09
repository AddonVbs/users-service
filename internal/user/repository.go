package user

import (
	"gorm.io/gorm"
)

// crud
type UsersRepository interface {
	CreateUser(user *User) error
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(user User) error
	DeleteUser(id int) error
	GetUserForTasksByRepo(userID int) (User, error)
}

type UserRepositoryDb struct {
	db *gorm.DB
}

// GetUserForTasksByRepo implements UsersRepository.
func (u *UserRepositoryDb) GetUserForTasksByRepo(userID int) (User, error) {
	var user User
	err := u.db.Preload("Tasks").First(&user, "id = ?", userID).Error
	return user, err
}

// CreateUser implements UsersRepository.
func (u *UserRepositoryDb) CreateUser(user *User) error {
	return u.db.Create(user).Error
}

// DeleteUser implements UsersRepository.
func (u *UserRepositoryDb) DeleteUser(id int) error {
	return u.db.Delete(&User{}, "id = ?", id).Error
}

// GetAllUser implements UsersRepository.
func (u *UserRepositoryDb) GetAllUser() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

// GetUser implements UsersRepository.
func (u *UserRepositoryDb) GetUser(id int) (User, error) {
	var user User
	err := u.db.First(&user, "id = ?", id).Error
	return user, err
}

// UpdataUser implements UsersRepository.
func (u *UserRepositoryDb) UpdataUser(user User) error {
	return u.db.Save(&user).Error
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &UserRepositoryDb{db: db}

}
