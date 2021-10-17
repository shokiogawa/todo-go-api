package controller

import (
	"example.com/go-mod-test/src/presentation"
	"example.com/go-mod-test/src/usecase/command"
)

type Controller struct {
	auth            *presentation.Auth
	saveUseCase     *command.SaveUseCase
	registerUseCase *command.RegisterUseCase
}

func NewController(
	auth *presentation.Auth,
	saveUseCase *command.SaveUseCase,
	registerUseCase *command.RegisterUseCase,
) *Controller {
	controller := new(Controller)
	controller.auth = auth
	controller.saveUseCase = saveUseCase
	controller.registerUseCase = registerUseCase
	return controller
}
