package models

type UserView struct {
	ID     uint
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func (r *User) ToUserView() UserView{
	return UserView{ID:r.ID, Name: r.Name, Email: r.Email}
}
