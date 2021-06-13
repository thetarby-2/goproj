package models

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
