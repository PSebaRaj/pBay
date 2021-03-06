package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique" required:"true"`
	Password []byte `json:"-" required:"true"`
}
