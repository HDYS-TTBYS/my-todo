package usecase_test

import (
	"testing"

	"HDYS-TTBYS/my-todo/domain/entities"
	mock_repository "HDYS-TTBYS/my-todo/mock/repository"
	"HDYS-TTBYS/my-todo/usecase"

	"github.com/golang/mock/gomock"
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
