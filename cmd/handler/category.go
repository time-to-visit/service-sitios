package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerCategory(e *echo.Echo, CategoryCaseUse usecase.CategoryUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	categoryEntry := entry.NewCategoryEntry(CategoryCaseUse)
	e.GET("/sites/category", categoryEntry.SeeCategory, auth)
	e.DELETE("/sites/category/:ID", categoryEntry.DeleteCategory, auth)
	e.POST("/sites/category", categoryEntry.InsertCategory, auth, validator.ValidateCategory)
	e.PUT("/sites/category", categoryEntry.UpdateCategory, auth, validator.ValidateCategory)
	return e
}
