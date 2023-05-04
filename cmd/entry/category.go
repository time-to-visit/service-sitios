package entry

import (
	"net/http"
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type categoryEntry struct {
	categoryCaseuse usecase.CategoryUseCase
}

func NewCategoryEntry(categoryCaseuse usecase.CategoryUseCase) *categoryEntry {
	return &categoryEntry{
		categoryCaseuse,
	}
}

func (e *categoryEntry) SeeCategory(c echo.Context) error {
	response, status := e.categoryCaseuse.SeeCategory()
	return c.JSON(status, response)
}

func (e *categoryEntry) UpdateCategory(c echo.Context) error {
	category := c.Get("category").(*entity.Category)
	response, status := e.categoryCaseuse.UpdateCategory(*category)
	return c.JSON(status, response)
}

func (e *categoryEntry) DeleteCategory(c echo.Context) error {
	id := c.Param("ID")
	idCategory, err := strconv.Atoi(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	response, status := e.categoryCaseuse.DeleteCategory(c.Request().Context(), idCategory)
	return c.JSON(status, response)
}

func (e *categoryEntry) InsertCategory(c echo.Context) error {
	category := c.Get("category").(*entity.Category)
	response, status := e.categoryCaseuse.InsertCategory(c.Request().Context(), *category)
	return c.JSON(status, response)
}
