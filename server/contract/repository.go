package contract

import "golang-tutorial/entity"

type Repository struct {
	User  UserRepository
	Todo  TodoRepository
}

// type exampleRepository interface {
// Code here
// }

type UserRepository interface {
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) error
	CheckEmail(email string) (bool, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type TodoRepository interface {
	CreateTodo(todo *entity.Todo) error
	GetTodo(id int) (*entity.Todo, error)
	UpdateTodo(id int, todo *entity.Todo) error
	DeleteTodo(id int) error
}
