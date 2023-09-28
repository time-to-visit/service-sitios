package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerDepartment(e *echo.Echo, DepartmentCaseUse usecase.DepartmentUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	departmentEntry := entry.NewDepartmentEntry(DepartmentCaseUse)
	e.GET("/sites/deparment", departmentEntry.SeeDeparment, auth)
	e.POST("/sites/deparment", departmentEntry.InsertDeparment, auth, validator.ValidateDepartment)
	e.DELETE("/sites/deparment/:ID", departmentEntry.DeleteDeparment, auth)
	return e
}
