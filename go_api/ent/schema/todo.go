package schema

import (
	"time"

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
		field.Time("created_at").
			Optional().
			Default(time.Now()),
		field.String("description").
			Optional(),
		field.Bool("is_complete").
			Default(false),
		field.String("title"),
		field.Time("updated_at").
			Optional().
			Default(time.Now()),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
