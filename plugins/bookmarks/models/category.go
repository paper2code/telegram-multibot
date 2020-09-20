package models

import (
	"gorm.io/gorm"
)

// Category represents a category
type Category struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}
