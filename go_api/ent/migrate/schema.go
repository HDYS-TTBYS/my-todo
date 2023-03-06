// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "assagin_person", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString},
		{Name: "is_complete", Type: field.TypeBool, Default: false},
		{Name: "title", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
	}
)

func init() {
}
