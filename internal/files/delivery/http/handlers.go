package http

import (
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

// Files handlers
type fileHandlers struct {
	cfg    *config.Config
	filesUC files.UseCase
	logger logger.Logger
}

// NewFileHandlers File handlers constructor
func NewFileHandlers(cfg *config.Config, filesUC files.UseCase, logger logger.Logger) files.Handlers {
	return &fileHandlers{cfg: cfg, filesUC: filesUC, logger: logger}
}
