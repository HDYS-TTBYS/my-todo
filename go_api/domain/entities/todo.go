package entities

type ResponseTodos struct {
	Total int `form:"total" json:"total"`
	ToDos []*ToDo
}
