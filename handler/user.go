package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sullyh7/devpage/view/home"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return Render(c, 200, home.Index())
}
