package files

// Files use case
type UseCase interface {
	Upload()
	Download()
	Delete()
	Share()
	Update()
}
