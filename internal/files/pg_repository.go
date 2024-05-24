package files

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/google/uuid"
)

// Files repository
type Repository interface {
	Upload(ctx context.Context, file *models.File) error
	Download(ctx context.Context, file *models.File) (*models.File, error)
	Delete(ctx context.Context, user_id, file uuid.UUID) error
	Share(ctx context.Context, share *models.Share) error
	GetAllFiles(ctx context.Context, user *models.User, pq *utils.PaginationQuery) (*models.FileList, error)
}
