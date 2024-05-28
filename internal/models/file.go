package models

import (
	"time"

	"github.com/google/uuid"
)

// File base model
type File struct {
	FileID    uuid.UUID `json:"file_id" db:"file_id" validate:"omitempty,uuid"`
	AuthorID  uuid.UUID `json:"author_id,omitempty" db:"author_id" validate:"required"`
	Title     string    `json:"title" db:"title" validate:"required,gte=1"`
	Content   []byte    `json:"content" db:"content" validate:"required"`
	Size      int64     `json:"size" db:"size"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type FilenameBase struct {
	FileID    uuid.UUID `json:"file_id" db:"files_id" validate:"omitempty,uuid"`
	AuthorID  uuid.UUID `json:"author_id,omitempty" db:"author_id" validate:"required"`
	Title     string    `json:"title" db:"title" validate:"required,gte=1"`
	Size      int64     `json:"size" db:"size"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// All Files response
type FileList struct {
	TotalCount int             `json:"total_count"`
	TotalPages int             `json:"total_pages"`
	Page       int             `json:"page"`
	Size       int             `json:"size"`
	HasMore    bool            `json:"has_more"`
	Files      []*FilenameBase `json:"files"`
}

// File base
type FileBase struct {
	FileID    uuid.UUID `json:"file_id" db:"file_id" validate:"omitempty,uuid"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id" validate:"omitempty,uuid"`
	Title     string    `json:"title" db:"title" validate:"required,gte=1"`
	Content   []byte    `json:"content" db:"content"`
	Size      int64     `json:"size" db:"size"`
	Author    string    `json:"author" db:"author"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Share struct {
	AuthorID uuid.UUID `json:"author_id" db:"author_id"`
	User_id  uuid.UUID `json:"user_id" db:"user_id"`
	File_id  uuid.UUID `json:"file_id" db:"files_id"`
}

type ShareResponse struct {
	FileID   uuid.UUID `json:"file_id" db:"files_id" validate:"omitempty,uuid"`
	AuthorID uuid.UUID `json:"author_id,omitempty" db:"author_id" validate:"required"`
	Title    string    `json:"title" db:"title" validate:"required,gte=1"`
	Size      int64     `json:"size" db:"size,omitempty"`
	Share     string    `db:"share"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
