package models

import "gorm.io/gorm"

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"uniqueIndex:udx_email"`
	Password  string         `json:"password"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:udx_email"`
}
