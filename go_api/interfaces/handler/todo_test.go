package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/interfaces/handler"
	mock_usecase "HDYS-TTBYS/my-todo/mock/usecase"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandler_FindMany(t *testing.T) {
	t.Run(
		"成功",
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
}
