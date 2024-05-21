package usecase

import (
	"context"
	"fmt"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

// Files UseCase
type filesUC struct {
	cfg       *config.Config
	filesRepo files.Repository
	logger    logger.Logger
}

// Files UseCase constructor
func NewFilesUseCase(cfg *config.Config, filesRepo files.Repository, logger logger.Logger) files.UseCase {
	fmt.Println("In NewFilesUseCase")
	return &filesUC{cfg: cfg, filesRepo: filesRepo, logger: logger}
}

// Upload file
func (u *filesUC) Upload(ctx context.Context) {
	fmt.Println("In (u *filesUC) Upload()")
	u.filesRepo.Upload(ctx)
}

// Download file
func (u *filesUC) Download() {
	fmt.Println("In (u *filesUC) Download()")

}

// Delete file
func (u *filesUC) Delete() {
	fmt.Println("In (u *filesUC) Delete()")

}

// Share file
func (u *filesUC) Share() {
	fmt.Println("In (u *filesUC) Share()")

}

// Update file
func (u *filesUC) Update() {
	fmt.Println("In (u *filesUC) Update()")

}
