package entity

import "github.com/google/uuid"

type User struct {
	Id           uint32
	PublicUserId string
	Name         string
	Email        string
	Password     string
}

func NewUser(id uint32, name string, email string, password string) *User {
	user := new(User)
	user.Id = id
	user.PublicUserId = uuid.New().String()
	user.Name = name
	user.Email = email
	user.Password = password
	return user
}
