package files

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
)

// Files use case
type UseCase interface {
	Upload(ctx context.Context, filename string, filesize int64, content *[]byte) error
	Download(ctx context.Context, file_id *[]byte) (*models.File, error)
	Delete(ctx context.Context, file_id *[]byte) error
	Share(ctx context.Context, user_id *[]byte) error
	Update()
}
