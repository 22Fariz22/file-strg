package files

import (
	"context"
	"mime/multipart"
)

// Files use case
type UseCase interface {
	Upload(ctx context.Context, file *multipart.FileHeader) error
	Download()
	Delete()
	Share()
	Update()
}
