package repository_implement

import (
	"example.com/go-mod-test/src/domain/entity"
	"example.com/go-mod-test/src/infrastructure"
)

type UserRepository struct {
	database infrastructure.Mysql
}

func NewUserRepository(database infrastructure.Mysql) *UserRepository {
	repo := new(UserRepository)
	repo.database = database
	return repo
}

func (repo *UserRepository) Register(user *entity.User) (publicUserId string, err error) {
	db, err := repo.database.Connect()
	if err != nil {
		return
	}
	query := `
INSERT INTO users (public_user_id, name, email, password)
VALUE (?,?,?,?)
`
	result := db.MustExec(query, user.PublicUserId, user.Name, user.Email, user.Password)
	num, err := result.RowsAffected()
	if num == 0 {
		return
	}
	publicUserId = user.PublicUserId
	return
}
