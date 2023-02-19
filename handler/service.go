package handler

import (
	"context"

	"github.com/eitarox/todo-app-api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password string, role entity.Role) (*entity.User, error)
}
