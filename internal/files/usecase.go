package files

import (
	"context"
)

// Files use case
type UseCase interface {
	Upload(ctx context.Context, filename string,filesize int64, content *[]byte) error
	Download()
	Delete()
	Share()
	Update()
}
