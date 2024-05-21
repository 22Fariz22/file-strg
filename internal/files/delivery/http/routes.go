package http

import (
	"github.com/AleksK1NG/api-mc/internal/files"
	"github.com/AleksK1NG/api-mc/internal/middleware"
	"github.com/labstack/echo/v4"
)

// Map news routes
func MapNewsRoutes(filesGroup *echo.Group, h files.Handlers, mw *middleware.MiddlewareManager) {
	// newsGroup.POST("/create", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.GET("/:news_id", h.GetByID())
	// newsGroup.GET("/search", h.SearchByTitle())
	// newsGroup.GET("", h.GetNews())
}
