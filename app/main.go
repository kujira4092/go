package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"

	"app/server"
)

func main(){
	echo := echo.New()

	echo.Use(middleware.Recover())
	echo.Use(middleware.Logger())

	server.InitRouter(echo)
	echo.Logger.Fatal(echo.Start(":8080"))
}
