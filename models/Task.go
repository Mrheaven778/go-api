package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required" `
	UserID      uint   `json:"user_id" validate:"required"`
}
