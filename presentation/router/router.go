package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	fmt.Println(v1)

	// v1.POST("/users", userCreate.NewController().Do)
}
