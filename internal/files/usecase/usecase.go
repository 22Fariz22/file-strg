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

	"github.com/google/uuid"
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
	return &filesUC{cfg: cfg, filesRepo: filesRepo, logger: logger}
}

// Upload file
func (u *filesUC) Upload(ctx context.Context, filename string, filesize int64, content *[]byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "filesUC.Upload")
	defer span.Finish()

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return httpErrors.NewUnauthorizedError(errors.WithMessage(err, "filesUC.Upload.GetUserFromCtx"))
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
func (u *filesUC) Download(ctx context.Context, fileIdBytes *[]byte) (*models.File, error) {
	fileIdUiid, err := uuid.ParseBytes(*fileIdBytes)
	if err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "filesUC.Download.ParseBytes"))
	}

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.WithMessage(err, "filesUC.Download.GetUserFromCtx"))
	}

	f := &models.File{}
	f.AuthorID = user.UserID
	f.FileID = fileIdUiid

	file, err := u.filesRepo.Download(ctx, f)
	if err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "filesUC.Download.filesRepo.Download"))
	}

	return file, nil
}

// Delete file
func (u *filesUC) Delete(ctx context.Context, fileIdBytes *[]byte) error {
	file_id, err := uuid.ParseBytes(*fileIdBytes)
	if err != nil {
		return httpErrors.NewBadRequestError(errors.WithMessage(err, "filesUC.Download.ParseBytes"))
	}

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		return httpErrors.NewUnauthorizedError(errors.WithMessage(err, "filesUC.Download.GetUserFromCtx"))
	}

	err = u.filesRepo.Delete(ctx, user.UserID, file_id)
	if err != nil {
		return err
	}

	return nil
}

// Share file
func (u *filesUC) Share(ctx context.Context, user_id *[]byte) error {
	fmt.Println("In (u *filesUC) Share()")

	return nil
}

// Update file
func (u *filesUC) Update() {
	fmt.Println("In (u *filesUC) Update()")

}
