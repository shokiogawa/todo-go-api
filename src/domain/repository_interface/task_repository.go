package repository_interface

import "example.com/go-mod-test/src/domain/entity"

type TaskRepository interface {
	Save(task *entity.Task) (err error)
}
