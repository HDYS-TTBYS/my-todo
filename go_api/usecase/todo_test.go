package usecase_test

import (
	"errors"
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

func returnTodo() *entities.ToDo {
	return &entities.ToDo{
		AssaginPerson: toPtr[string]("hdys"),
		CreatedAt:     toPtr[int64](1),
		Description:   toPtr[string]("description"),
		Id:            toPtr[int](1),
		IsComplete:    toPtr[bool](false),
		Title:         "title",
		UpdatedAt:     toPtr[int64](1),
	}
}

func postTodoJsonBody() *entities.PostTodoJSONBody {
	return &entities.PostTodoJSONBody{
		AssiginPerson: "hdys",
		Description:   toPtr[string]("description"),
		Title:         "title",
	}
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
			mockRepo.EXPECT().FindMany(0).Return(
				[]*entities.ToDo{returnTodo()}, nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todos, err := u.FindMany(0)
			if assert.NoError(tt, err) {
				assert.Equal(tt,
					&entities.ResponseTodos{
						Total: 1,
						ToDos: []*entities.ToDo{returnTodo()},
					}, todos)
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
			mockRepo.EXPECT().FindById(1).Return(nil, errors.New("error"))
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
			mockRepo.EXPECT().FindById(1).Return(returnTodo(), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.FindById(1)
			if assert.NoError(tt, err) {
				assert.Equal(tt, returnTodo(), todo)
			}
		},
	)
}

func TestTodoUseCase_Create(t *testing.T) {
	t.Run(
		"成功",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Create(postTodoJsonBody()).Return(returnTodo(), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Create(postTodoJsonBody())
			if assert.NoError(tt, err) {
				assert.Equal(tt, returnTodo(), todo)
			}
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Create(postTodoJsonBody()).Return(nil, errors.New("error"))
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Create(postTodoJsonBody())
			if assert.Error(tt, err) {
				assert.Nil(tt, todo)
			}
		},
	)
}
