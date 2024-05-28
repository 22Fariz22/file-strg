package usecase

// import (
// 	"context"
// 	"testing"

// 	"github.com/AleksK1NG/api-mc/internal/files/mock"
// 	"github.com/AleksK1NG/api-mc/internal/models"
// 	"github.com/AleksK1NG/api-mc/pkg/logger"
// 	"github.com/AleksK1NG/api-mc/pkg/utils"
// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/opentracing/opentracing-go"
// )

// func TestFilesUC_Upload(t *testing.T){
// 	t.Parallel()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	apiLogger := logger.NewApiLogger(nil)
// 	mockFilesRepo := mock.NewMockRepository(ctrl)
// 	filesUC := NewFilesUseCase(nil, mockFilesRepo, apiLogger)

// 	userUID := uuid.New()

// 	content := make([]byte,10)

// 	file := &models.File{
// 	AuthorID: userUID,
// 	Title:    "Title long text string greater then 1 characters",
// 	Content:  content,
//   }

// 	user := &models.User{
// 		UserID: userUID,
// 	}

// 	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
// 	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "filesUC.Create")
// 	defer span.Finish()

// 		// mockNewsRepo.EXPECT().Create(ctxWithTrace, gomock.Eq(news)).Return(news, nil)
  
// 		mockFilesRepo.EXPECT().Upload(ctxWithTrace,gomock.Eq(file)).Return(nil)
	
// }