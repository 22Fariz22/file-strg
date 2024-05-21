package files

// Files repository
type Repository interface {
	Upload()
	Download()
	Delete()
	Share()
	Update()
}
