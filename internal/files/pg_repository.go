package files

import "context"

// Files repository
type Repository interface {
	Upload(ctx context.Context)
	Download()
	Delete()
	Share()
	Update()
}
