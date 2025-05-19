package repository

import (
	"golang-tutorial/contract"
	"golang-tutorial/entity"

	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func implTodoRepository(db *gorm.DB) contract.TodoRepository {
	return &TodoRepo{
		db: db,
	}
}

func (r *TodoRepo) GetTodo(id int) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.Table("todo").Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, err
}

func (r *TodoRepo) CreateTodo(todo *entity.Todo) error {
	return r.db.Table("todo").Create(todo).Error
}

func (r *TodoRepo) UpdateTodo(id int, todo *entity.Todo) error {
	return r.db.Where("id = ?", id).Updates(&todo).Error
}

func (r *TodoRepo) DeleteTodo(id int) error {
	return r.db.Where("id = ?", id).Delete(&entity.Todo{}).Error
}
