package repository_implement

import (
	"example.com/go-mod-test/src/domain/entity"
	"example.com/go-mod-test/src/infrastructure"
)

type TaskRepository struct {
	database infrastructure.Mysql
}

func NewTaskRepository(database infrastructure.Mysql) *TaskRepository {
	repo := new(TaskRepository)
	repo.database = database
	return repo
}

func (repo *TaskRepository) Save(task *entity.Task) (err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	query := `
INSERT INTO tasks (public_task_id,user_id,title, content)
VALUE (?,?,?,?)
`
	result := db.MustExec(query, task.PublicUserId, task.UserId, task.Title, task.Content)
	num, err := result.RowsAffected()
	if num == 0 {
		return
	}
	return
}
