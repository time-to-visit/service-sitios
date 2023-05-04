package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerSubCategory(e *echo.Echo, subCategoryCaseUse usecase.SubCategoryUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	subCategoryEntry := entry.NewSubcategoryEntry(subCategoryCaseUse)
	e.POST("/subcategory", subCategoryEntry.InsertSubcategory, auth, validator.ValidateSubCategory)
	e.DELETE("/subcategory/:IDCATEGORY/:IDSUBCATEGORY", subCategoryEntry.DeleteCategory, auth)
	return e
}
