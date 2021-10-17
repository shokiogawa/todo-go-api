package main

import (
	"example.com/go-mod-test/src/presentation/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func RouterSetting(controller *controller.Controller) (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//register
	e.POST("/register", controller.Register)
	//task
	e.POST("/task", controller.SaveTask)
	//user
	u := e.Group("/users")
	u.Use(middleware.JWT([]byte(os.Getenv("SIGNINGKEY"))))
	u.GET("/user", controller.GetUser)
	return e
}
