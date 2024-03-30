package router

import (
	"github.com/labstack/echo/v4"
	_ "swclabs/swipe-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// Docs
// API documentation
// Router documentation
// Base on: http://${HOST}:${PORT}/docs/index.html#/
func Docs(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)
}
