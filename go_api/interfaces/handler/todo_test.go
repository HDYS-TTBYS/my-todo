package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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
		"no id",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/todo/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.FindByID(c))
		},
	)
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

func TestTodoHandler_Create(t *testing.T) {
	t.Run(
		"no body",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/todo", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Create(c))
		},
	)
	t.Run(
		"bad body",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.PostTodoJsonBodyBad())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Create(c))
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().Create(utils.PostTodoJsonBody()).Return(nil, errors.New("error"))
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.PostTodoJsonBody())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Create(c))
		},
	)
	t.Run(
		"成功",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().Create(utils.PostTodoJsonBody()).Return(utils.ReturnTodo(), nil)
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.PostTodoJsonBody())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			if assert.NoError(tt, h.Create(c)) {
				assert.Equal(tt, http.StatusCreated, rec.Code)
				assert.Equal(tt, fmt.Sprintln(`{"assagin_person":"hdys","created_at":1,"description":"description","id":1,"is_complete":false,"title":"title","updated_at":1}`), rec.Body.String())
			}
		},
	)
}

func TestTodoHandler_Update(t *testing.T) {
	t.Run(
		"no id",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.UpdateTodoJsonBody())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPatch, "/api/todo/:id", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Update(c))
		},
	)
	t.Run(
		"bad body",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.UpdateTodoJsonBodyBad())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPatch, "/api/todo/:id", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Update(c))
		},
	)
	t.Run(
		"失敗",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().Update(utils.UpdateTodoJsonBody(), 1).Return(nil, errors.New("error"))
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.UpdateTodoJsonBody())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPatch, "/api/todo/:id", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")
			h := handler.NewTodoHandler(mockUsecase)
			assert.Error(tt, h.Update(c))
		},
	)
	t.Run(
		"成功",
		func(tt *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsecase := mock_usecase.NewMockITodoUseCase(mockCtrl)
			mockUsecase.EXPECT().Update(utils.UpdateTodoJsonBody(), 1).Return(utils.ReturnTodo(), nil)
			e := echo.New()
			bodyBytes, err := json.Marshal(utils.UpdateTodoJsonBody())
			if err != nil {
				panic(err)
			}
			req := httptest.NewRequest(http.MethodPatch, "/api/todo/:id", strings.NewReader(string(bodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")
			h := handler.NewTodoHandler(mockUsecase)
			if assert.NoError(tt, h.Update(c)) {
				assert.Equal(tt, http.StatusOK, rec.Code)
				assert.Equal(tt, fmt.Sprintln(`{"assagin_person":"hdys","created_at":1,"description":"description","id":1,"is_complete":false,"title":"title","updated_at":1}`), rec.Body.String())
			}
		},
	)
}
