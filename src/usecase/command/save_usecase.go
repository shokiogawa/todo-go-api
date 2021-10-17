package command

import (
	"example.com/go-mod-test/src/domain/entity"
	"example.com/go-mod-test/src/domain/repository_interface"
)

type SaveUseCase struct {
	repository repository_interface.TaskRepository
}

func NewSaveUseCase(repository repository_interface.TaskRepository) *SaveUseCase {
	uc := new(SaveUseCase)
	uc.repository = repository
	return uc
}

func (uc *SaveUseCase) Invoke(title string, content string, userId uint32) (err error) {
	task := entity.NewTask(0, title, content, userId)
	err = uc.repository.Save(task)
	return
}
