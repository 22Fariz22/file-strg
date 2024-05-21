package repository

import (
	"context"
	"fmt"

	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/jmoiron/sqlx"
)

// Files Repository
type filesRepo struct {
	db *sqlx.DB
}

// Files repository constructor
func NewFilesRepository(db *sqlx.DB) files.Repository {
	fmt.Println("In NewFilesRepository()")
	return &filesRepo{db: db}
}

// Upload file
func (r *filesRepo) Upload(ctx context.Context) {
	fmt.Println("In (r *filesRepo) Upload() ")
}

// Download file
func (r *filesRepo) Download() {
	fmt.Println("In (r *filesRepo) Download() ")
}

// Delete file
func (r *filesRepo) Delete() {
	fmt.Println("In (r *filesRepo) Delete() ")
}

// Share file
func (r *filesRepo) Share() {
	fmt.Println("In (r *filesRepo) Share() ")
}

// Update file
func (r *filesRepo) Update() {
	fmt.Println("In (r *filesRepo) Update() ")
}
