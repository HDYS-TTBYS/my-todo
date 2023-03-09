package usecase

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/domain/repository"
)

// Todoのusecase
type ITodoUseCase interface {
	FindMany(offset int) (*entities.ResponseTodos, error)
	FindById(id int) (*entities.ToDo, error)
	Create(todo *entities.PostTodoJSONBody) (*entities.ToDo, error)
	Update(todo *entities.UpdateTodoIdJSONBody, id int) (*entities.ToDo, error)
	Delete(id int) error
}

type todoUseCase struct {
	todoRepository repository.ITodoRepository
}

func NewTodoUseCase(tr repository.ITodoRepository) ITodoUseCase {
	return &todoUseCase{todoRepository: tr}
}

func (tu todoUseCase) FindMany(offset int) (todo *entities.ResponseTodos, err error) {
	count, err := tu.todoRepository.TotalCount()
	if err != nil {
		return nil, err
	}
	if *count == 0 {
		return &entities.ResponseTodos{Total: *count, ToDos: []*entities.ToDo{}}, nil
	}
	todos, err := tu.todoRepository.FindMany(offset)
	if err != nil {
		return nil, err
	}
	return &entities.ResponseTodos{Total: *count, ToDos: todos}, nil
}

func (tu todoUseCase) FindById(id int) (todo *entities.ToDo, err error) {
	rtodo, err := tu.todoRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return rtodo, err
}

func (tu todoUseCase) Create(todo *entities.PostTodoJSONBody) (*entities.ToDo, error) {
	rtodo, err := tu.todoRepository.Create(todo)
	if err != nil {
		return nil, err
	}
	return rtodo, err
}

func (tu todoUseCase) Update(todo *entities.UpdateTodoIdJSONBody, id int) (*entities.ToDo, error) {
	rtodo, err := tu.todoRepository.Update(todo, id)
	if err != nil {
		return nil, err
	}
	return rtodo, err
}

func (tu todoUseCase) Delete(id int) error {
	err := tu.todoRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
