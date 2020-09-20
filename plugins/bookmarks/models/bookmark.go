package models

import (
	"gorm.io/gorm"
)

// Bookmark represents a bookmark
type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
	Tags Tags   `json:"tags" gorm:"type:string[]"`
}
