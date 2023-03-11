package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/interfaces/handler"
	mock_usecase "HDYS-TTBYS/my-todo/mock/usecase"
	"HDYS-TTBYS/my-todo/utils"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandler_FindMany(t *testing.T) {
	t.Run(
		"offsetなし",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().FindMany(0).Return(
				&entities.ResponseTodos{
					Total: 0,
					ToDos: []*entities.ToDo{}},
				nil,
			)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			if assert.NoError(tt, h.FindMany(c)) {
				assert.Equal(tt, http.StatusOK, rec.Code)
				assert.Equal(tt, fmt.Sprintln(`{"total":0,"ToDos":[]}`), rec.Body.String())
			}
		},
	)
	t.Run(
		"offsetあり",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().FindMany(0).Return(
				&entities.ResponseTodos{
					Total: 0,
					ToDos: []*entities.ToDo{}},
				nil,
			)
			e := echo.New()
			q := make(url.Values)
			q.Set("offset", "0")
			req := httptest.NewRequest(http.MethodGet, "/api/todos?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			if assert.NoError(tt, h.FindMany(c)) {
				assert.Equal(tt, http.StatusOK, rec.Code)
				assert.Equal(tt, fmt.Sprintln(`{"total":0,"ToDos":[]}`), rec.Body.String())
			}
		},
	)
}

func TestTodoHandler_FindByID(t *testing.T) {
	t.Run(
		"データあり",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().FindById(1).Return(utils.ReturnTodo(), nil)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/todo/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")
			h := handler.NewTodoHandler(mockUsecase)
			if assert.NoError(tt, h.FindByID(c)) {
				assert.Equal(tt, http.StatusOK, rec.Code)
				assert.Equal(tt, fmt.Sprintln(`{"assagin_person":"hdys","created_at":1,"description":"description","id":1,"is_complete":false,"title":"title","updated_at":1}`), rec.Body.String())
			}
		},
	)
	t.Run(
		"データなし",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().FindById(1111).Return(nil, errors.New("error"))
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/todo/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1111")
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.FindByID(c))
		},
	)
}
