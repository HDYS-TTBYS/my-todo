package repository

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/domain/repository"
	"HDYS-TTBYS/my-todo/ent"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type todoRepository struct {
	ec  *ent.Client
	ctx context.Context
}

type TodoOmitDesc struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AssaginPerson holds the value of the "assagin_person" field.
	AssaginPerson string `json:"assagin_person,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// IsComplete holds the value of the "is_complete" field.
	IsComplete bool `json:"is_complete,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewTodoRepository(ec *ent.Client, ctx context.Context) repository.ITodoRepository {
	return &todoRepository{ec, ctx}
}

// ent.Todo -> entities.ToDo
func dataTransform(t *ent.Todo) *entities.ToDo {
	ca := t.CreatedAt.Unix()
	ua := t.UpdatedAt.Unix()
	return &entities.ToDo{
		AssaginPerson: &t.AssaginPerson,
		CreatedAt:     &ca,
		Description:   &t.Description,
		Id:            &t.ID,
		IsComplete:    &t.IsComplete,
		Title:         t.Title,
		UpdatedAt:     &ua}
}

func dataTransformF(t *TodoOmitDesc) *entities.ToDo {
	ca := t.CreatedAt.Unix()
	ua := t.UpdatedAt.Unix()
	return &entities.ToDo{
		AssaginPerson: &t.AssaginPerson,
		CreatedAt:     &ca,
		Description:   nil,
		Id:            &t.ID,
		IsComplete:    &t.IsComplete,
		Title:         t.Title,
		UpdatedAt:     &ua}
}

func (tr *todoRepository) TotalCount() (*int, error) {
	total, err := tr.ec.Todo.Query().Count(tr.ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed querying total count todo")
	}
	return &total, nil
}

func (tr *todoRepository) FindMany(offset int) ([]*entities.ToDo, error) {
	var t []TodoOmitDesc
	err := tr.ec.Todo.Query().
		Limit(20).
		Offset(offset).
		Order(ent.Desc("created_at")).
		Select("assagin_person", "created_at", "is_complete", "title", "updated_at").
		Scan(tr.ctx, &t)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed querying todos")
	}
	var todos []*entities.ToDo
	for _, v := range t {
		todos = append(todos, dataTransformF(&v))
	}
	return todos, err
}

func (tr *todoRepository) FindById(id int) (*entities.ToDo, error) {
	t, err := tr.ec.Todo.Get(tr.ctx, id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "failed querying todo")
	}
	rt := dataTransform(t)
	return rt, nil
}

func (tr *todoRepository) Create(todo *entities.PostTodoJSONBody) (*entities.ToDo, error) {
	t, err := tr.ec.Todo.Create().
		SetAssaginPerson(todo.AssiginPerson).
		SetDescription(*todo.Description).
		SetTitle(todo.Title).
		Save(tr.ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed creating todo")
	}
	rt := dataTransform(t)
	return rt, nil
}

func (tr *todoRepository) Update(todo *entities.UpdateTodoIdJSONBody, id int) (*entities.ToDo, error) {
	t, err := tr.ec.Todo.UpdateOneID(id).
		SetAssaginPerson(todo.AssiginPerson).
		SetDescription(*todo.Description).
		SetIsComplete(todo.IsComplete).SetTitle(todo.Title).
		SetUpdatedAt(time.Now()).
		Save(tr.ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "failed updating todo")
	}
	rt := dataTransform(t)
	return rt, nil
}

func (tr *todoRepository) Delete(id int) error {
	err := tr.ec.Todo.DeleteOneID(id).Exec(tr.ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "failed deleting todo")
	}
	return nil
}
