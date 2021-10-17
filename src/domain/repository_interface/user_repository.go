package repository_interface

import "example.com/go-mod-test/src/domain/entity"

type UserRepository interface {
	Register(user *entity.User) (publicUserId string, err error)
}
