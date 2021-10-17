package command

import (
	"example.com/go-mod-test/src/domain/entity"
	"example.com/go-mod-test/src/domain/repository_interface"
)

type RegisterUseCase struct {
	repo repository_interface.UserRepository
}

func NewRegisterUseCase(repo repository_interface.UserRepository) *RegisterUseCase {
	uc := new(RegisterUseCase)
	uc.repo = repo
	return uc
}

func (uc *RegisterUseCase) Invoke(name string, email string, password string) (publicUserId string, err error) {
	user := entity.NewUser(0, name, email, password)
	publicUserId, err = uc.repo.Register(user)
	if err != nil {
		return
	}
	return
}

//type User struct {
//	Id           uint32
//	PublicUserId string
//	Name         string
//	Email        string
//	Password     string
//}
