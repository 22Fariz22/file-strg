package http

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

// Files handlers
type fileHandlers struct {
	cfg     *config.Config
	filesUC files.UseCase
	logger  logger.Logger
}

// NewFileHandlers File handlers constructor
func NewFileHandlers(cfg *config.Config, filesUC files.UseCase, logger logger.Logger) files.Handlers {
	fmt.Println("In NewFileHandlers()")
	return &fileHandlers{cfg: cfg, filesUC: filesUC, logger: logger}
}

func (h fileHandlers) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "newsHandlers.Create")
		defer span.Finish()

		h.filesUC.Upload(ctx)
		return c.JSON(http.StatusCreated, nil)

	}
}

func (h fileHandlers) Download() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers) Share() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}
