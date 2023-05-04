package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type deparmentEntry struct {
	deparmentCaseuse usecase.DepartmentUseCase
}

func NewDepartmentEntry(deparmentCaseuse usecase.DepartmentUseCase) *deparmentEntry {
	return &deparmentEntry{
		deparmentCaseuse,
	}
}

func (e *deparmentEntry) SeeDeparment(c echo.Context) error {
	response, status := e.deparmentCaseuse.SeeDeparment()
	return c.JSON(status, response)
}

func (e *deparmentEntry) InsertDeparment(c echo.Context) error {
	department := c.Get("department").(*entity.Department)
	response, status := e.deparmentCaseuse.InsertDeparment(c.Request().Context(), *department)
	return c.JSON(status, response)
}

func (e *deparmentEntry) DeleteDeparment(c echo.Context) error {
	id := c.Param("ID")
	idDepartment, _ := strconv.Atoi(id)
	response, status := e.deparmentCaseuse.DeleteDeparment(c.Request().Context(), idDepartment)
	return c.JSON(status, response)
}
