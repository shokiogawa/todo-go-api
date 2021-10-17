package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserInfoJson struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (con *Controller) Register(c echo.Context) (err error) {
	var jsonUser UserInfoJson
	err = json.NewDecoder(c.Request().Body).Decode(&jsonUser)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	publicUserId, err := con.registerUseCase.Invoke(jsonUser.Name, jsonUser.Email, jsonUser.Password)
	tokenString, err := con.auth.GetToken(publicUserId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	c.Response().Write([]byte(tokenString))
	return
}
