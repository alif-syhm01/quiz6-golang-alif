package models

// User model struct
type User struct {
	Id       uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
}
