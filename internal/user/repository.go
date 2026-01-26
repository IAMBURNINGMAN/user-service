package user

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

func (r *userRepository) CreateUser(user *UserStruct) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]UserStruct, error) {
	var users []UserStruct
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id uint) (UserStruct, error) {
	var thatUser UserStruct
	err := r.db.First(&thatUser, id).Error
	return thatUser, err
}

func (r *userRepository) UpdateUser(user *UserStruct) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&UserStruct{}, id).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
