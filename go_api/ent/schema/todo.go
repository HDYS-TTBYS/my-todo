package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("assagin_person"),
		field.Time("created_at"),
		field.String("description").
			Nillable(),
		field.Bool("is_complete").
			Default(false),
		field.String("title"),
		field.Time("updated_at"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
