package utils

import "HDYS-TTBYS/my-todo/domain/entities"

func ToPtr[T any](s T) *T {
	return &s
}

func ReturnTodo() *entities.ToDo {
	return &entities.ToDo{
		AssaginPerson: ToPtr[string]("hdys"),
		CreatedAt:     ToPtr[int64](1),
		Description:   ToPtr[string]("description"),
		Id:            ToPtr[int](1),
		IsComplete:    ToPtr[bool](false),
		Title:         "title",
		UpdatedAt:     ToPtr[int64](1),
	}
}

func PostTodoJsonBody() *entities.PostTodoJSONBody {
	return &entities.PostTodoJSONBody{
		AssiginPerson: "hdys",
		Description:   ToPtr[string]("description"),
		Title:         "title",
	}
}

func PostTodoJsonBodyBad() *entities.PostTodoJSONBody {
	return &entities.PostTodoJSONBody{
		AssiginPerson: "",
		Description:   ToPtr[string](""),
		Title:         "",
	}
}

func UpdateTodoJsonBody() *entities.UpdateTodoIdJSONBody {
	return &entities.UpdateTodoIdJSONBody{
		AssiginPerson: "hdys",
		Description:   ToPtr[string]("description"),
		IsComplete:    false,
		Title:         "title",
	}
}
