package files

import "github.com/labstack/echo/v4"

// Files HTTP Handlers interface
type Handlers interface {
	Upload() echo.HandlerFunc
	Download() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Share() echo.HandlerFunc
	GetAllFiles() echo.HandlerFunc
	Update() echo.HandlerFunc
}
