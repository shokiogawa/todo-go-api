package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskInfoJson struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

func (con *Controller) SaveTask(c echo.Context) error {
	var jsonTask TaskInfoJson
	err := json.NewDecoder(c.Request().Body).Decode(&jsonTask)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	err = con.saveUseCase.Invoke(jsonTask.Title, jsonTask.Content, uint32(jsonTask.UserId))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "ok")
}
