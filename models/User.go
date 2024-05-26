package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Tasks    []Task `json:"tasks"`
}
