package repository

import (
	"HDYS-TTBYS/my-todo/domain/entities"
)

// Todoのrepository
type ITodoRepository interface {
	TotalCount() (*int, error)
	FindMany(offset int) ([]*entities.ToDo, error)
	FindById(id int) (*entities.ToDo, error)
	Create(todo *entities.PostTodoJSONBody) (*entities.ToDo, error)
	Update(todo *entities.UpdateTodoIdJSONBody, id int) (*entities.ToDo, error)
	Delete(id int) error
}
