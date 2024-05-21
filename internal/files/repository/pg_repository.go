package repository

import (
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/jmoiron/sqlx"
)

// Files Repository
type filesRepo struct {
	db *sqlx.DB
}

// Files repository constructor
func NewFilesRepository(db *sqlx.DB) files.Repository {
	return &filesRepo{db: db}
}
