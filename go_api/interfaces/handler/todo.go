package handler

import (
	"HDYS-TTBYS/my-todo/domain/entities"
	"HDYS-TTBYS/my-todo/interfaces/models"
	"HDYS-TTBYS/my-todo/usecase"
	"net/http"
	"strconv"

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
	stringOffset := c.QueryParam("offset")
	intOffset, err := strconv.Atoi(stringOffset)
	if err != nil {
		intOffset = 0
	}
	todos, err := h.ITodoUseCase.FindMany(intOffset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *todoHandler) FindByID(c echo.Context) error {
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	todo, err := h.ITodoUseCase.FindById(intId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *todoHandler) Create(c echo.Context) error {
	var ptodo models.PostTodoJSONBody
	if err := c.Bind(&ptodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	if err := ptodo.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	ntodo, err := h.ITodoUseCase.Create((*entities.PostTodoJSONBody)(&ptodo))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ntodo)
}

func (h *todoHandler) Update(c echo.Context) error {
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	var utodo models.UpdateTodoIdJSONBody
	if err := c.Bind(&utodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	if err := utodo.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	rtodo, err := h.ITodoUseCase.Update((*entities.UpdateTodoIdJSONBody)(&utodo), intId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rtodo)
}

func (h *todoHandler) Delete(c echo.Context) error {
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	err = h.ITodoUseCase.Delete(intId)
	if err != nil {
		return err
	}
	tmp := "contant deleted"
	p := &tmp
	return c.JSON(http.StatusOK, models.Message{Message: p})
}
