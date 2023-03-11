package usecase_test

import (
	"errors"
	"testing"

	"HDYS-TTBYS/my-todo/domain/entities"
	mock_repository "HDYS-TTBYS/my-todo/mock/repository"
	"HDYS-TTBYS/my-todo/usecase"
	"HDYS-TTBYS/my-todo/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTodoUseCase_FindMany(t *testing.T) {
	t.Run(
		"データなし",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().TotalCount().Return(utils.ToPtr[int](0), nil)
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
			mockRepo.EXPECT().TotalCount().Return(utils.ToPtr[int](1), nil)
			mockRepo.EXPECT().FindMany(0).Return(
				[]*entities.ToDo{utils.ReturnTodo()}, nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todos, err := u.FindMany(0)
			if assert.NoError(tt, err) {
				assert.Equal(tt,
					&entities.ResponseTodos{
						Total: 1,
						ToDos: []*entities.ToDo{utils.ReturnTodo()},
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
			mockRepo.EXPECT().FindById(1).Return(utils.ReturnTodo(), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.FindById(1)
			if assert.NoError(tt, err) {
				assert.Equal(tt, utils.ReturnTodo(), todo)
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
			mockRepo.EXPECT().Create(utils.RostTodoJsonBody()).Return(utils.ReturnTodo(), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Create(utils.RostTodoJsonBody())
			if assert.NoError(tt, err) {
				assert.Equal(tt, utils.ReturnTodo(), todo)
			}
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Create(utils.RostTodoJsonBody()).Return(nil, errors.New("error"))
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Create(utils.RostTodoJsonBody())
			if assert.Error(tt, err) {
				assert.Nil(tt, todo)
			}
		},
	)
}

func TestTodoUseCase_Update(t *testing.T) {
	t.Run(
		"成功",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Update(utils.UpdateTodoJsonBody(), 1).Return(utils.ReturnTodo(), nil)
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Update(utils.UpdateTodoJsonBody(), 1)
			if assert.NoError(tt, err) {
				assert.Equal(tt, utils.ReturnTodo(), todo)
			}
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Update(utils.UpdateTodoJsonBody(), 1).Return(nil, errors.New("error"))
			u := usecase.NewTodoUseCase(mockRepo)
			todo, err := u.Update(utils.UpdateTodoJsonBody(), 1)
			if assert.Error(tt, err) {
				assert.Nil(tt, todo)
			}
		},
	)
}

func TestTodoUseCase_Delete(t *testing.T) {
	t.Run(
		"成功",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Delete(1).Return(nil)
			u := usecase.NewTodoUseCase(mockRepo)
			err := u.Delete(1)
			assert.NoError(tt, err)
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock_repository.NewMockITodoRepository(mockCtrl)
			mockRepo.EXPECT().Delete(1).Return(errors.New("error"))
			u := usecase.NewTodoUseCase(mockRepo)
			err := u.Delete(1)
			assert.Error(tt, err)
		},
	)
}
