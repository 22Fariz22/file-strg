package models

import (
	"time"

	"github.com/google/uuid"
)

// File base model
type File struct {
	FileID    uuid.UUID `json:"file_id" db:"file_id" validate:"omitempty,uuid"`
	AuthorID  uuid.UUID `json:"author_id,omitempty" db:"author_id" validate:"required"`
	Title     string    `json:"title" db:"title" validate:"required,gte=10"`
	Content   []byte    `json:"content" db:"content"`
	Size      int64     `json:"size" db:"size"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// All Files response
type FileList struct {
	TotalCount int      `json:"total_count"`
	TotalPages int      `json:"total_pages"`
	Page       int      `json:"page"`
	Size       int      `json:"size"`
	HasMore    bool     `json:"has_more"`
	Files      []string `json:"files"`
}

// File base
type FileBase struct {
	FileID    uuid.UUID `json:"file_id" db:"file_id" validate:"omitempty,uuid"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id" validate:"omitempty,uuid"`
	Title     string    `json:"title" db:"title" validate:"required,gte=10"`
	Content   []byte    `json:"content" db:"content"`
	Size      int64     `json:"size" db:"size"`
	Author    string    `json:"author" db:"author"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
