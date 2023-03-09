package repository

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/ent/enttest"
	"context"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestTodoRepository_TotalCount(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := NewTodoRepository(client, c)
			count, err := r.TotalCount()
			if assert.NoError(tt, err) {
				zero := 0
				p := &zero
				assert.Equal(tt, count, p)
			}
		},
	)
	t.Run(
		"データ1個",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := NewTodoRepository(client, c)
			now := time.Now()
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetCreatedAt(now).
				SetDescription("test description").
				SetTitle("test title").
				SetUpdatedAt(now).
				Save(c)
			if assert.NoError(ttt, err) {
				count, err := r.TotalCount()
				if assert.NoError(ttt, err) {
					one := 1
					p := &one
					assert.Equal(ttt, count, p)
				}
			}
		},
	)
}

func TestTodoRepository_FindMany(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := NewTodoRepository(client, c)
			todos, err := r.FindMany(0)
			if assert.NoError(tt, err) {
				var t []*entities.ToDo
				assert.Equal(tt, todos, t)
			}
		},
	)
	t.Run(
		"データ1個",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := NewTodoRepository(client, c)
			now := time.Now()
			createdTodo, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetCreatedAt(now).
				SetDescription("test description").
				SetTitle("test title").
				SetUpdatedAt(now).
				Save(c)
			if assert.NoError(ttt, err) {
				todos, err := r.FindMany(0)
				if assert.NoError(ttt, err) {
					var ttodo []*entities.ToDo
					ttodo = append(ttodo, dataTransform(createdTodo))
					assert.Equal(ttt, todos, ttodo)
				}
			}
		},
	)
}
