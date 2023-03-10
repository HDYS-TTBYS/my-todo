package usecase_test

import (
	"net/http"
	"testing"

	"HDYS-TTBYS/my-todo/domain/entities"
	mock_repository "HDYS-TTBYS/my-todo/mock/repository"
	"HDYS-TTBYS/my-todo/usecase"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func toPtr[T any](s T) *T {
	return &s
}

func TestTodoUseCase_FindMany(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().TotalCount().Return(toPtr[int](0), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todos, err := u.FindMany(0)
			if assert.NoError(tt, err) {
				assert.Equal(tt, &entities.ResponseTodos{Total: 0, ToDos: []*entities.ToDo{}}, todos)
			}
		},
	)
	t.Run(
		"データ1個あり",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().TotalCount().Return(toPtr[int](1), nil)
			a := &entities.ToDo{
				AssaginPerson: toPtr[string]("hdys"),
				CreatedAt:     toPtr[int64](1),
				Description:   toPtr[string]("description"),
				Id:            toPtr[int](1),
				IsComplete:    toPtr[bool](false),
				Title:         "title",
				UpdatedAt:     toPtr[int64](1),
			}
			mockRepo.EXPECT().FindMany(0).Return(
				[]*entities.ToDo{a}, nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todos, err := u.FindMany(0)
			if assert.NoError(tt, err) {
				assert.Equal(tt, &entities.ResponseTodos{Total: 1, ToDos: []*entities.ToDo{a}}, todos)
			}
		},
	)
}

func TestTodoUseCase_FindById(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().FindById(1).Return(
				nil,
				echo.NewHTTPError(http.StatusInternalServerError, ""))
			u := usecase.NewTodoUseCase(mockRepo)
			_, err := u.FindById(1)
			assert.Error(tt, err)
		},
	)
	t.Run(
		"データ1個あり",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			a := &entities.ToDo{
				AssaginPerson: toPtr[string]("hdys"),
				CreatedAt:     toPtr[int64](1),
				Description:   toPtr[string]("description"),
				Id:            toPtr[int](1),
				IsComplete:    toPtr[bool](false),
				Title:         "title",
				UpdatedAt:     toPtr[int64](1),
			}
			mockRepo.EXPECT().FindById(1).Return(a, nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.FindById(1)
			if assert.NoError(tt, err) {
				assert.Equal(tt, a, todo)
			}
		},
	)
}
