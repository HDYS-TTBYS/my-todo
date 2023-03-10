package repository_test

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/ent/enttest"
	"HDYS-TTBYS/my-todo/infrastructure/repository"
	"context"
	"testing"

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
			r := repository.NewTodoRepository(client, c)
			count, err := r.TotalCount()
			if assert.NoError(tt, err) {
				assert.Equal(tt, 0, *count)
			}
		},
	)
	t.Run(
		"データ1個",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetDescription("test description").
				SetTitle("test title").
				Save(c)
			if assert.NoError(ttt, err) {
				count, err := r.TotalCount()
				if assert.NoError(ttt, err) {
					assert.Equal(ttt, 1, *count)
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
			r := repository.NewTodoRepository(client, c)
			todos, err := r.FindMany(0)
			if assert.NoError(tt, err) {
				var expected []*entities.ToDo
				assert.Equal(tt, expected, todos)
			}
		},
	)
	t.Run(
		"データ1個",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetDescription("test description").
				SetTitle("test title").
				Save(c)
			if assert.NoError(ttt, err) {
				todos, err := r.FindMany(0)
				if assert.NoError(ttt, err) {
					assert.Equal(ttt, "hdys", *todos[0].AssaginPerson)
					assert.Equal(ttt, "test description", *todos[0].Description)
					assert.Equal(ttt, "test title", todos[0].Title)
				}
			}
		},
	)
}

func TestTodoRepository_FindById(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			todo, err := r.FindById(1)
			if assert.Error(tt, err) {
				assert.Nil(tt, todo)
			}
		},
	)
	t.Run(
		"データ1個",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetDescription("test description").
				SetTitle("test title").
				Save(c)
			if assert.NoError(ttt, err) {
				todo, err := r.FindById(1)
				if assert.NoError(ttt, err) {
					assert.Equal(ttt, "hdys", *todo.AssaginPerson)
					assert.Equal(ttt, "test description", *todo.Description)
					assert.Equal(ttt, "test title", todo.Title)
				}
			}
		},
	)
}

func TestTodoRepository_Create(t *testing.T) {
	t.Run(
		"成功",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			desc := "test description"
			postTodoJsonBody := &entities.PostTodoJSONBody{
				AssiginPerson: "hdys",
				Description:   &desc,
				Title:         "test title"}
			createdTodo, err := r.Create(postTodoJsonBody)
			if assert.NoError(tt, err) {
				assert.Equal(tt, postTodoJsonBody.AssiginPerson, *createdTodo.AssaginPerson)
				assert.Equal(tt, *postTodoJsonBody.Description, *createdTodo.Description)
				assert.Equal(tt, postTodoJsonBody.Title, createdTodo.Title)
			}
		},
	)
}

func TestTodoRepository_Update(t *testing.T) {
	t.Run(
		"データなし失敗",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			desc := "test description update"
			updataTodoJsonBody := &entities.UpdateTodoIdJSONBody{
				AssiginPerson: "hdys",
				Description:   &desc,
				IsComplete:    true,
				Title:         "test title update",
			}
			_, err := r.Update(updataTodoJsonBody, 1)
			assert.Error(tt, err)
		},
	)
	t.Run(
		"成功",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetDescription("test description").
				SetTitle("test title").
				Save(c)
			if assert.NoError(ttt, err) {
				desc := "test description update"
				updataTodoJsonBody := &entities.UpdateTodoIdJSONBody{
					AssiginPerson: "hdys",
					Description:   &desc,
					IsComplete:    true,
					Title:         "test title update",
				}
				updatedTodo, err := r.Update(updataTodoJsonBody, 1)
				if assert.NoError(ttt, err) {
					assert.Equal(ttt, updataTodoJsonBody.AssiginPerson, *updatedTodo.AssaginPerson)
					assert.Equal(ttt, *updataTodoJsonBody.Description, *updatedTodo.Description)
					assert.Equal(ttt, updataTodoJsonBody.IsComplete, *updatedTodo.IsComplete)
					assert.Equal(ttt, updataTodoJsonBody.Title, updatedTodo.Title)
				}
			}
		},
	)
}

func TestTodoRepository_Delete(t *testing.T) {
	t.Run(
		"データなし失敗",
		func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			err := r.Delete(1)
			assert.Error(tt, err)
		},
	)
	t.Run(
		"成功",
		func(ttt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			c := context.Background()
			r := repository.NewTodoRepository(client, c)
			_, err := client.Todo.Create().
				SetAssaginPerson("hdys").
				SetDescription("test description").
				SetTitle("test title").
				Save(c)
			if assert.NoError(ttt, err) {
				err := r.Delete(1)
				if assert.NoError(ttt, err) {
					count, err := client.Todo.Query().Count(c)
					if assert.NoError(ttt, err) {
						assert.Equal(ttt, 0, count)
					}
				}
			}
		},
	)
}
