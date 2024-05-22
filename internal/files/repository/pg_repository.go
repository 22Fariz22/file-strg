package repository

import (
	"context"
	"fmt"

	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
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
func (r *filesRepo) Upload(ctx context.Context, file *models.File) error {
	fmt.Println("IN (r *filesRepo) Upload() ")
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesRepo.Upload")
	defer span.Finish()

	res, err := r.db.ExecContext(ctx, "INSERT INTO files (author_id, title, content, size) VALUES($1,$2,$3,$4)",
		file.AuthorID, file.Title, file.Content, file.Size)
	if err != nil {
		errors.Wrap(err, "filesRepo.Upload.Exec")
		return err
	}
	fmt.Println("RES: ", res)

	return nil
}

// Download file
func (r *filesRepo) Download(ctx context.Context, file *models.File) (*models.File, error) {
	fmt.Println("In (r *filesRepo) Download() ")
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.Download")
	defer span.Finish()

	if err := r.db.QueryRowxContext(ctx,
		`SELECT title, content FROM files WHERE author_id = $1 and files_id = $2`,
		file.AuthorID, file.FileID).StructScan(file); err != nil {
		return nil, errors.Wrap(err, "fileRepo.Download.QueryRowxContext")
	}

	return file, nil
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
