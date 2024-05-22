package http

import (
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/middleware"
	"github.com/labstack/echo/v4"
)

// Map files routes
func MapFilesRoutes(filesGroup *echo.Group, h files.Handlers, mw *middleware.MiddlewareManager) {
	filesGroup.Use(mw.AuthSessionMiddleware)
	filesGroup.POST("/upload", h.Upload(), mw.AuthSessionMiddleware)
	filesGroup.GET("/download/", h.Download(), mw.AuthSessionMiddleware)
	filesGroup.DELETE("/delete", h.Delete(), mw.AuthSessionMiddleware)
	filesGroup.POST("/share/:file_id", h.Share(), mw.AuthSessionMiddleware, mw.CSRF)
	filesGroup.PUT("/update/:file_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
}
