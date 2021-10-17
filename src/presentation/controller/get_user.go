package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (con *Controller) GetUser(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	publicUserId := claims["sub"].(string)
	fmt.Println(publicUserId)

	return c.String(http.StatusOK, "ok")
}
