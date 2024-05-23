package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/google/uuid"
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
func (r *filesRepo) Delete(ctx context.Context, user_id, file_id uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx,
		`DELETE FROM files WHERE author_id = $1 and files_id = $2`,
		user_id, file_id)
	if err != nil {
		return errors.Wrap(err, "filesRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "filesRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "filesRepo.Delete.rowsAffected")
	}

	return nil
}

// Share file
func (r *filesRepo) Share(ctx context.Context, share *models.Share) error {
	fmt.Println("In (r *filesRepo) Share() ")
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.Share")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, `UPDATE files SET share = $1, updated_at = CURRENT_TIMESTAMP WHERE author_id = $2 AND files_id = $3`,
		share.User_id, share.AuthorID, share.File_id)
	if err != nil {
		return errors.Wrap(err, "filesRepo.Share.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "filesRepo.Share.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "filesRepo.Share.rowsAffected")
	}

	return nil
}

func (r *filesRepo)	GetAllFiles(ctx context.Context,user *models.User,pq *utils.PaginationQuery) (*models.FileList, error){
	fmt.Println("In (r *filesRepo) GetAllFiles() ")
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.GetAllFiles")
	defer span.Finish()

	getTotalCount := `SELECT COUNT(files_id) FROM files WHERE author_id = $1`

  var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalCount,user.UserID); err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetFiles.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.FileList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Files:       make([]*models.FilenameBase, 0),
		}, nil
	}

		getFiles := `SELECT files_id, author_id, title, size, updated_at, created_at 
				FROM files 
				WHERE author_id = $1
				ORDER BY created_at, updated_at OFFSET $2 LIMIT $3`

	var filesList = make([]*models.FilenameBase, 0, pq.GetSize())
	rows, err := r.db.QueryxContext(ctx, getFiles,user.UserID, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetFiles.QueryxContext")
	}
	defer rows.Close()

  for rows.Next() {
		n := &models.FilenameBase{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "filesRepo.GetAllFiles.StructScan")
		}
		filesList = append(filesList, n)
	}
  if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "filesRepo.GetFiles.rows.Err")
	}

	return &models.FileList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Files:       filesList,
	}, nil

}


// Update file
func (r *filesRepo) Update() {
	fmt.Println("In (r *filesRepo) Update() ")
}
