package http

import (
	"fmt"

	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Files handlers
type fileHandlers struct {
	cfg    *config.Config
	filesUC files.UseCase
	logger logger.Logger
}

// NewFileHandlers File handlers constructor
func NewFileHandlers(cfg *config.Config, filesUC files.UseCase, logger logger.Logger) files.Handlers {
	fmt.Println("In NewFileHandlers()")
	return &fileHandlers{cfg: cfg, filesUC: filesUC, logger: logger}
}

func (h fileHandlers)Upload()echo.HandlerFunc{
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers)Download()echo.HandlerFunc{
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers)Delete()echo.HandlerFunc{
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers)Share()echo.HandlerFunc{
	return func(ctx echo.Context) error {
		return nil
	}
}

func (h fileHandlers)Update()echo.HandlerFunc{
	return func(ctx echo.Context) error {
		return nil
	}
}

