package signup

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() Controller {
	return Controller{}
}

func (c Controller) Do(ctx echo.Context) error {
	fmt.Println("Hello, SignUp!")
	return nil
}
