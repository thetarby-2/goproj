package db

import (
	"goproj2/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(id uint) (models.User, error)                         // fetches user by its id
	GetAll() ([]models.User, error)                               // fetches all users and returns an array of users
	Update(toUpdate *models.User, updateInput *models.User) error // updates User passed in 'toUpdate' parameter to have fields as in 'updateInput' parameter
	Delete(user *models.User) error
	Create(user *models.User) error
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetById(id uint) (models.User, error) {
	var user models.User
	res := r.DB.First(&user, id).Error
	return user, res
}

func (r *UserRepository) Update(toUpdate *models.User, update *models.User) error {
	err := r.DB.Model(toUpdate).Updates(models.User{Name: update.Name, Password: update.Password}).Error
	return err
}

func (r *UserRepository) Create(user *models.User) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *UserRepository) Delete(user *models.User) error {
	err := r.DB.Delete(&user).Error
	return err
}
