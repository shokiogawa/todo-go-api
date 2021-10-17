package entity

import "github.com/google/uuid"

type Task struct {
	Id           uint32
	PublicUserId string
	Title        string
	Content      string
	UserId       uint32
}

func NewTask(id uint32, title string, content string, userId uint32) *Task {
	task := new(Task)
	task.Id = id
	task.PublicUserId = uuid.New().String()
	task.Title = title
	task.Content = content
	task.UserId = userId
	return task
}
