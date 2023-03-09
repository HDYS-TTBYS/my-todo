package models

import (
	"HDYS-TTBYS/my-todo/domain/entities"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetTodosParams struct {
	Offset int `query:"offset"`
}

type Message entities.Error

type PostTodoJSONBody entities.PostTodoJSONBody

type UpdateTodoIdJSONBody entities.UpdateTodoIdJSONBody

func (g GetTodosParams) Validate() error {
	return validation.Validate(&g.Offset,
		validation.NotNil,
		validation.Min(0),
	)
}

func (p PostTodoJSONBody) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.AssiginPerson, validation.Required, validation.Length(1, 20)),
		validation.Field(&p.Description, validation.Length(0, 255)),
		validation.Field(&p.Title, validation.Required, validation.Length(1, 128)),
	)
}

func (u UpdateTodoIdJSONBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.AssiginPerson, validation.Required, validation.Min(1), validation.Max(20)),
		validation.Field(&u.Description, validation.Max(255)),
		validation.Field(&u.Title, validation.Required, validation.Min(1), validation.Max(128)),
		validation.Field(&u.IsComplete, validation.NotNil),
	)
}
