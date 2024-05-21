package http

import (
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/middleware"
	"github.com/labstack/echo/v4"
)

// Map files routes
func MapFilesRoutes(filesGroup *echo.Group, h files.Handlers, mw *middleware.MiddlewareManager) {
	filesGroup.POST("/upload",h.Upload(),mw.AuthSessionMiddleware,mw.CSRF)
	filesGroup.POST("/download/:file_id",h.Download(),mw.AuthSessionMiddleware,mw.CSRF)
	filesGroup.DELETE("/delete/:file_id",h.Delete(),mw.AuthSessionMiddleware,mw.CSRF)
	filesGroup.POST("/share/:file_id",h.Share(),mw.AuthSessionMiddleware,mw.CSRF)
	filesGroup.PUT("/update/:file_id",h.Update(),mw.AuthSessionMiddleware,mw.CSRF)
}