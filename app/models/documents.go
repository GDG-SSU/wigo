package models

import "github.com/jinzhu/gorm"

/**
 * Document - A database model representing Wiki document
 */
type Document struct {
	gorm.Model
	Title   string `gorm:"size:255"sql:"not null;"`
	Content string `sql:"type:text;not null;"`
}
