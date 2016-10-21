package models

import "time"

/**
 * Document - A database model representing Wiki document
 */
type Document struct {
    ID          uint
    Title       string `gorm:"size:255"sql:"not null;"`
    Content     string `sql:"type:text;not null;"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   *time.Time
}
