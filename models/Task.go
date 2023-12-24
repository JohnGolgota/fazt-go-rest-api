package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null;unique_index" json:"title"` // Add json tag
	Description string `json:"description"`                        // Add json tag
	Completed   bool   `gorm:"default:false" json:"completed"`     // Add json tag
	UserID      uint   `json:"user_id"`                            // Add json tag
}
