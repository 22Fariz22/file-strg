package files

import (
	"context"

	"github.com/AleksK1NG/api-mc/internal/models"
)

// Files repository
type Repository interface {
	Upload(ctx context.Context, file *models.File)error
	Download()
	Delete()
	Share()
	Update()
}
