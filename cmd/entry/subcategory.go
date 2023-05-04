package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type subcategoryEntry struct {
	subcategoryCaseuse usecase.SubCategoryUseCase
}

func NewSubcategoryEntry(subcategoryCaseuse usecase.SubCategoryUseCase) *subcategoryEntry {
	return &subcategoryEntry{
		subcategoryCaseuse,
	}
}

func (e *subcategoryEntry) InsertSubcategory(c echo.Context) error {
	subCategory := c.Get("subCategory").(*entity.Subcategory)
	response, status := e.subcategoryCaseuse.InsertSubcategory(c.Request().Context(), *subCategory)
	return c.JSON(status, response)
}

func (e *subcategoryEntry) DeleteCategory(c echo.Context) error {
	idSubCategory, _ := strconv.Atoi(c.Param("IDSUBCATEGORY"))
	idCategory, _ := strconv.Atoi(c.Param("IDCATEGORY"))
	response, status := e.subcategoryCaseuse.DeleteSubCategory(c.Request().Context(), idCategory, idSubCategory)
	return c.JSON(status, response)
}
