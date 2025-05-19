package service

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"golang-tutorial/entity"
	"net/http"
)

type TodoService struct {
	todoRepository contract.TodoRepository
}

func implTodoService(repo *contract.Repository) contract.TodoService {
	return &TodoService{
		todoRepository: repo.Todo,
	}
}

func (s *TodoService) GetTodo(todoID int) (*dto.TodoResponse, error) {
	todo, err := s.todoRepository.GetTodo(todoID)
	if err != nil {
		return nil, err
	}

	response := &dto.TodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data",
		Data: dto.TodoData{
			ID:            todo.ID,
			Todo:          todo.Todo,
		},
	}
	return response, nil
}

func (s *TodoService) CreateTodo(payload *dto.TodoRequest) (*dto.TodoResponse, error) {
	todo := &entity.Todo{
		Todo:          payload.Todo,
	}

	err := s.todoRepository.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	response := &dto.TodoResponse{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil membuat data",
		Data: dto.TodoData{
			ID:            todo.ID,
			Todo:          todo.Todo,
		},
	}

	return response, nil
}

func (s *TodoService) UpdateTodo(id int, payload *dto.TodoRequest) (*dto.TodoResponse, error) {
	todo := &entity.Todo{
		Todo:          payload.Todo,
	}

	err := s.todoRepository.UpdateTodo(id, todo)
	if err != nil {
		return nil, err
	}

	response := &dto.TodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mengubah data",
		Data: dto.TodoData{
			ID:            todo.ID,
			Todo:          todo.Todo,
		},
	}

	return response, nil
}

func (s *TodoService) DeleteTodo(id int) (*dto.TodoResponse, error) {
	err := s.todoRepository.DeleteTodo(id)
	if err != nil {
		return nil, err
	}

	response := &dto.TodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil menghapus data",
	}

	return response, nil
}
