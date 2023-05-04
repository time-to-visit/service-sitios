package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerCategory(e *echo.Echo, CategoryCaseUse usecase.CategoryUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	categoryEntry := entry.NewCategoryEntry(CategoryCaseUse)
	e.GET("/category", categoryEntry.SeeCategory, auth)
	e.DELETE("category/:ID", categoryEntry.DeleteCategory, auth)
	e.POST("/category", categoryEntry.InsertCategory, auth, validator.ValidateCategory)
	e.PUT("/category", categoryEntry.UpdateCategory, auth, validator.ValidateCategory)
	return e
}
