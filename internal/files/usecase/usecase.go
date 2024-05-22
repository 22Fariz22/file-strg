package usecase

import (
	"context"
	"fmt"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/AleksK1NG/api-mc/pkg/utils"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
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
func (u *filesUC) Upload(ctx context.Context, filename string,filesize int64, content *[]byte) error {
	fmt.Println("In (u *filesUC) Upload()")
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.Upload")
	defer span.Finish()

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return httpErrors.NewUnauthorizedError(errors.WithMessage(err, "filesUC.Upl.GetUserFromCtx"))
	}

	f := &models.File{}
	f.AuthorID = user.UserID
	f.Title = filename
	f.Content = *content
	f.Size = filesize

	if err = utils.ValidateStruct(ctx, f); err != nil {
		return httpErrors.NewBadRequestError(errors.WithMessage(err, "newsUC.Create.ValidateStruct"))
	}

	u.filesRepo.Upload(ctx, f)
	return nil
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
