package main

import (
	"embed"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/devpage/handler"
)

//go:embed public
var fs embed.FS

func main() {
	app := echo.New()
	app.StaticFS("/*", fs)
	userHandler := handler.UserHandler{}

	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")
	fmt.Println("it is working")
}
