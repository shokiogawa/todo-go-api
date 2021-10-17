package main

import (
	"example.com/go-mod-test/src/infrastructure"
	"example.com/go-mod-test/src/infrastructure/repository_implement"
	"example.com/go-mod-test/src/presentation"
	controller2 "example.com/go-mod-test/src/presentation/controller"
	"example.com/go-mod-test/src/usecase/command"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Initialize struct {
	Db         *sqlx.DB
	Controller *controller2.Controller
}

func NewInitialize() (init *Initialize, err error) {
	init = new(Initialize)
	db, err := infrastructure.NewMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	init.Db = db.DB

	taskRepo := repository_implement.NewTaskRepository(*db)
	registerRepo := repository_implement.NewUserRepository(*db)

	saveUseCase := command.NewSaveUseCase(taskRepo)
	registerUseCase := command.NewRegisterUseCase(registerRepo)

	auth := &presentation.Auth{}
	controller := controller2.NewController(auth, saveUseCase, registerUseCase)

	init.Controller = controller
	return
}
