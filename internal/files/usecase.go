package files

import "context"

// Files use case
type UseCase interface {
	Upload(ctx context.Context)
	Download()
	Delete()
	Share()
	Update()
}
