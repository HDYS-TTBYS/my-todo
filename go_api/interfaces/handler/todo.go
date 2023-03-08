package handler

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/interfaces/models"
	"HDYS-TTBYS/my-todo/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// echo.contextからパラメータを取り出す
// パラメータをvalidationする
// パラメータをusecaseに渡す
// 結果をbodyに書き込んで返す

type ITodoHandler interface {
	FindMany(c echo.Context) error
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type todoHandler struct {
	usecase.ITodoUseCase
}

func NewTodoHandler(uc usecase.ITodoUseCase) ITodoHandler {
	return &todoHandler{uc}
}

func (h *todoHandler) FindMany(c echo.Context) error {
	var offset models.GetTodosParams
	if err := c.Bind(&offset); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding offset: %w", err))
	}
	if err := offset.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG offset: %w", err))
	}
	todos, err := h.ITodoUseCase.FindMany(offset.Offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *todoHandler) FindByID(c echo.Context) error {
	var id models.GetTodoIdParam
	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding id: %w", err))
	}
	if err := id.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG offset: %w", err))
	}
	todo, err := h.ITodoUseCase.FindById(id.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *todoHandler) Create(c echo.Context) error {
	var ptodo models.PostTodoJSONBody
	if err := c.Bind(&ptodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding post todo: %w", err))
	}
	if err := ptodo.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG post todo: %w", err))
	}
	ntodo, err := h.ITodoUseCase.Create((*entities.PostTodoJSONBody)(&ptodo))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ntodo)
}

func (h *todoHandler) Update(c echo.Context) error {
	var id models.GetTodoIdParam
	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding id: %w", err))
	}
	if err := id.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG offset: %w", err))
	}
	var utodo models.UpdateTodoIdJSONBody
	if err := c.Bind(&utodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding patch todo: %w", err))
	}
	if err := utodo.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG patch todo: %w", err))
	}
	rtodo, err := h.ITodoUseCase.Update((*entities.UpdateTodoIdJSONBody)(&utodo), id.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rtodo)
}

func (h *todoHandler) Delete(c echo.Context) error {
	var id models.GetTodoIdParam
	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed binding id: %w", err))
	}
	if err := id.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("validate NG offset: %w", err))
	}
	err := h.ITodoUseCase.Delete(id.ID)
	if err != nil {
		return err
	}
	tmp := "contant deleted"
	p := &tmp
	return c.JSON(http.StatusOK, models.Message{Message: p})
}
