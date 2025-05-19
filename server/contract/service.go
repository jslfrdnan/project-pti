package contract

import "golang-tutorial/dto"

type Service struct {
	User  UserService
	Todo  TodoService
}

// type exampleService interface {
// Code here
// }


type UserService interface {
	GetUser(userID int) (*dto.UserResponse, error)
	Register(payload *dto.UserRequest) (*dto.UserResponse, error)
	Login(payload *dto.UserRequest) (*dto.UserResponse, error)
}

type TodoService interface {
	GetTodo(todoID int) (*dto.TodoResponse, error)
	CreateTodo(payload *dto.TodoRequest) (*dto.TodoResponse, error)
	UpdateTodo(id int, payload *dto.TodoRequest) (*dto.TodoResponse, error)
	DeleteTodo(id int) (*dto.TodoResponse, error)
}
