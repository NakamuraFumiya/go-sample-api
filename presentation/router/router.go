package router

import (
	"github.com/fumiyanakamura/go-sample-api/presentation/internalapi/v1/auth/signup"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/signup", signup.NewController().Do)
}
