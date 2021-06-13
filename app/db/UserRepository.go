package db

import (
	"goproj2/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]models.User, error)
	GetById(id uint, user interface{}) error
	Update(toUpdate *models.User, update *models.User) error
	Delete(user *models.User) error
	Create(user *models.User) error
}

type UserRepository struct{
	DB *gorm.DB
}


func (r *UserRepository) GetAll() ([]models.User, error){
	var users []models.User
	res := r.DB.Find(&users)
	return users, res.Error
}


func (r *UserRepository) GetById(id uint, user interface{}) error{
	res := r.DB.First(user, id).Error
	return res
}


func (r *UserRepository) Update(toUpdate *models.User, update *models.User) error {
	err := r.DB.Model(toUpdate).Updates(models.User{Name:update.Name, Password: update.Password}).Error
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